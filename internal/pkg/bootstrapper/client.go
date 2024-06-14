package bootstrapper

import (
	"context"
	"errors"
	"fmt"
	"time"

	resourceutils "github.com/nukleros/operator-builder-tools/pkg/resources"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	waitTimeout  = 120 * time.Second
	waitInterval = 5 * time.Second
)

var (
	ErrCreateKubeClient            = errors.New("unable to create kubernetes client")
	ErrCreateDiscoveryClient       = errors.New("unable to create discovery client")
	ErrGetAPIResources             = errors.New("unable to retrieve the api resources")
	ErrGetRESTMapping              = errors.New("unable to retrieve rest mapping")
	ErrMissingResource             = errors.New("unable to use nil resource")
	ErrMissingGroupVersionResource = errors.New("unable to use nil group version resource")
)

// Client represents the boostrapper client which is used to interact with the Kubernetes API for
// bootstrapping resources
type Client struct {
	Context context.Context
	Client  *dynamic.DynamicClient
	Config  *rest.Config
	API     *restmapper.DeferredDiscoveryRESTMapper
}

// NewClient returns a new instance of a client object.
func NewClient(ctx context.Context, kubeconfigFile string) (client *Client, err error) {
	var kubeconfig *rest.Config

	if kubeconfigFile != "" {
		kubeconfig, err = clientcmd.BuildConfigFromFlags("", kubeconfigFile)
		if err != nil {
			return nil, fmt.Errorf("error creating kubernetes client from file - %s", ErrCreateKubeClient.Error())
		}
	} else {
		kubeconfig, err = rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("error creating kubernetes client from in cluster config - %s", ErrCreateKubeClient.Error())
		}
	}

	// create a discovery client use for mapping GVK to GVR
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%s - %w", ErrCreateDiscoveryClient, err)
	}

	// create a rest mapper that is used to convert GVK to GVR
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(discoveryClient))

	// create the kubernetes dynamic client
	dynamicClient, err := dynamic.NewForConfig(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("error creating dynamic client - %w", err)
	}

	return &Client{
		Context: ctx,
		Config:  kubeconfig,
		Client:  dynamicClient,
		API:     mapper,
	}, nil
}

// Apply applies a particular resource to a cluster.
func (client *Client) Apply(resource *unstructured.Unstructured, gvr *schema.GroupVersionResource) error {
	if resource == nil {
		return ErrMissingResource
	}

	if gvr == nil {
		return ErrMissingGroupVersionResource
	}

	_, err := client.Client.Resource(*gvr).Namespace(resource.GetNamespace()).Apply(
		client.Context,
		resource.GetName(),
		resource,
		metav1.ApplyOptions{FieldManager: FieldManager},
	)
	if err != nil {
		return fmt.Errorf("unable to create resource - %w", err)
	}

	return nil
}

// Destroy removes a particular resource from a cluster.
func (client *Client) Destroy(resource *unstructured.Unstructured, gvr *schema.GroupVersionResource) error {
	if resource == nil {
		return ErrMissingResource
	}

	if gvr == nil {
		return ErrMissingGroupVersionResource
	}

	err := client.Client.Resource(*gvr).Namespace(resource.GetNamespace()).Delete(
		client.Context,
		resource.GetName(),
		metav1.DeleteOptions{},
	)
	if err != nil {
		return fmt.Errorf("unable to destroy resource - %w", err)
	}

	return nil
}

// WaitForReady waits for a particular resource to be ready.
func (client *Client) WaitForReady(resource *unstructured.Unstructured, gvr *schema.GroupVersionResource) error {
	if err := client.waitFor(resource, gvr, isReady); err != nil {
		return fmt.Errorf("unable to wait for resource readiness - %w", err)
	}

	return nil
}

// WaitForMissing waits for a particular resource to be missing.
func (client *Client) WaitForDestroy(resource *unstructured.Unstructured, gvr *schema.GroupVersionResource) error {
	if err := client.waitFor(resource, gvr, isMissing); err != nil {
		return fmt.Errorf("unable to wait for resource destruction - %w", err)
	}

	return nil
}

// GetGroupVersionResource returns the GroupVersionResource object from a particular resource.
func (client *Client) GetGroupVersionResource(gvk schema.GroupVersionKind) (*schema.GroupVersionResource, error) {
	mapping, err := client.API.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return nil, fmt.Errorf("%s - %w", ErrGetRESTMapping, err)
	}

	return &mapping.Resource, nil
}

// waitFor waits for a particular resource to condition to occur.
func (client *Client) waitFor(resource *unstructured.Unstructured, gvr *schema.GroupVersionResource, f checkFunction) error {
	if resource == nil {
		return ErrMissingResource
	}

	if gvr == nil {
		return ErrMissingGroupVersionResource
	}

	// run the initial check so we do not have to wait for the interval ticker
	condition, err := client.runCheck(resource, gvr, f)
	if err != nil {
		return fmt.Errorf("unable to determine resource condition - %w", err)
	}

	if condition {
		return nil
	}

	timeout, interval := time.After(waitTimeout), time.NewTicker(waitInterval)
	defer interval.Stop()

	for {
		select {
		case <-timeout:
			return fmt.Errorf("timed out waiting for resource")
		case <-interval.C:
			condition, err := client.runCheck(resource, gvr, f)
			if err != nil {
				return fmt.Errorf("unable to determine resource condition - %w", err)
			}

			if condition {
				return nil
			}
		}
	}
}

func (client *Client) runCheck(resource *unstructured.Unstructured, gvr *schema.GroupVersionResource, f checkFunction) (bool, error) {
	resourceFromCluster, err := client.Client.Resource(*gvr).Namespace(resource.GetNamespace()).Get(
		client.Context,
		resource.GetName(),
		metav1.GetOptions{},
	)

	if err != nil {
		if !apierrors.IsNotFound(err) {
			return false, fmt.Errorf("unable to get resource - %w", err)
		}

		// ensure we have a nil resource
		resourceFromCluster = nil
	}

	return f(resourceFromCluster)
}

type checkFunction func(*unstructured.Unstructured) (bool, error)

// isReady satisfies the checkFunction signature determining whether a given resource is ready or not.
func isReady(resource *unstructured.Unstructured) (bool, error) {
	if resource == nil {
		return false, nil
	}

	ready, err := resourceutils.IsReady(resource)
	if err != nil {
		return false, fmt.Errorf("unable to determine resource readiness - %w", err)
	}

	return ready, nil
}

// isMissing satisfies the checkFunction signature determining whether a given resource is missing or not.
func isMissing(resource *unstructured.Unstructured) (bool, error) {
	return (resource == nil), nil
}
