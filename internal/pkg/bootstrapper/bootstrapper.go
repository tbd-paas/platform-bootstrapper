package bootstrapper

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	FieldManager = "platform-bootstrapper"
)

// Bootstrapper represents the object that is performing the bootstrapping process.
type Bootstrapper struct {
	Context context.Context
	Client  *Client
	Log     zerolog.Logger
}

// NewBootstrapper returns a new instance of a bootstrapper object.
func NewBootstrapper(client *Client, options ...Option) *Bootstrapper {
	// determine the log level for the bootstrapper
	logLevel := zerolog.InfoLevel
	if HasOption(options, WithDebug) {
		logLevel = zerolog.DebugLevel
	}

	return &Bootstrapper{
		Context: client.Context,
		Client:  client,
		Log:     zerolog.New(os.Stdout).With().Timestamp().Logger().Level(logLevel),
	}
}

// RunAction runs a specific bootstrap action against a set of resources.
func (b *Bootstrapper) RunAction(action bootstrapAction, resources ...*unstructured.Unstructured) error {
	index := 0
	timeout := time.After(waitTimeout)

	// run the action against the resources
	for index < len(resources) {
		select {
		case <-timeout:
			return fmt.Errorf("timed out waiting for action completion")
		default:
			// get the group version resource for this resource
			gvr, err := b.Client.GetGroupVersionResource(resources[index].GroupVersionKind())
			if err != nil {
				proper := errors.Unwrap(err)

				// restart the loop if we have a NoKindMatchError.
				// TODO: we need to handle this better as we progress this project.
				//nolint:errorLint
				switch proper.(type) {
				case *meta.NoKindMatchError:
					b.Client.API.Reset()

					continue
				default:
					return fmt.Errorf("unable to get group version resource for: [%s] - %w", ResourceMessage(resources[index]), err)
				}
			}

			// run the specific action against this resource
			if err := action(resources[index], gvr); err != nil {
				return err
			}

			index++
		}
	}

	return nil
}

type bootstrapAction func(*unstructured.Unstructured, *schema.GroupVersionResource) error

// Apply runs the specific apply action for the bootstrapper.
func (b *Bootstrapper) Apply(resource *unstructured.Unstructured, gvr *schema.GroupVersionResource) error {
	// apply the resource to the cluster
	b.Log.Info().Msgf("creating resource - %s", ResourceMessage(resource))
	if err := b.Client.Apply(resource, gvr); err != nil {
		return fmt.Errorf("unable to apply resource: [%s] - %w", ResourceMessage(resource), err)
	}

	// wait for resource to be ready in the cluster before moving on
	b.Log.Info().Msgf("waiting for resource to be ready - %s", ResourceMessage(resource))
	if err := b.Client.WaitForReady(resource, gvr); err != nil {
		return fmt.Errorf("unable to wait for resource: [%s] - %w", ResourceMessage(resource), err)
	}

	return nil
}

// Destroy runs the specific destroy action for the bootstrapper.
func (b *Bootstrapper) Destroy(resource *unstructured.Unstructured, gvr *schema.GroupVersionResource) error {
	// remove the resource from the cluster
	b.Log.Info().Msgf("destroying resource - %s", ResourceMessage(resource))
	if err := b.Client.Destroy(resource, gvr); err != nil {
		return fmt.Errorf("unable to destroy resource: [%s] - %w", ResourceMessage(resource), err)
	}

	// wait for resource to be missing from the cluster before moving on
	b.Log.Info().Msgf("waiting for resource to be missing - %s", ResourceMessage(resource))
	if err := b.Client.WaitForDestroy(resource, gvr); err != nil {
		return fmt.Errorf("unable to wait for resource: [%s] - %w", ResourceMessage(resource), err)
	}

	return nil
}

// ResourceMessage returns the string used for logging messages for resources.
func ResourceMessage(resource *unstructured.Unstructured) string {
	return fmt.Sprintf("group=%s, version=%s, kind=%s, namespace=%s, name=%s",
		resource.GetObjectKind().GroupVersionKind().Group,
		resource.GetObjectKind().GroupVersionKind().Version,
		resource.GetObjectKind().GroupVersionKind().Kind,
		resource.GetNamespace(),
		resource.GetName(),
	)
}
