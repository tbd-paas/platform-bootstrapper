package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	resourceutils "github.com/nukleros/operator-builder-tools/pkg/resources"
	"github.com/rs/zerolog"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/tbd-paas/platform-bootstrapper/internal/pkg/resources"
)

const (
	waitTimeout  = 120 * time.Second
	waitInterval = 5 * time.Second
)

var (
	ErrCreateKubeClient = errors.New("unable to create kubernetes client")
)

func main() {
	// create the logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.InfoLevel)

	// create the kubernetes client
	var kubeconfig *rest.Config

	var err error

	kubeconfigFile := os.Getenv("KUBECONFIG")
	if kubeconfigFile != "" {
		kubeconfig, err = clientcmd.BuildConfigFromFlags("", kubeconfigFile)
		if err != nil {
			logger.Error().Msgf("error creating kubernetes client from file - %s", ErrCreateKubeClient.Error())
			panic(err)
		}
	} else {
		kubeconfig, err = rest.InClusterConfig()
		if err != nil {
			logger.Error().Msgf("error creating kubernetes client from in cluster config - %s", ErrCreateKubeClient.Error())
			panic(err)
		}
	}

	kubeClient, err := dynamic.NewForConfig(kubeconfig)
	if err != nil {
		logger.Error().Msgf("error creating kubernetes client rest config - %s", ErrCreateKubeClient.Error())
		panic(err)
	}

	// create a discovery client use for mapping GVK to GVR
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(kubeconfig)
	if err != nil {
		logger.Error().Msgf("error creating discovery client")
		panic(err)
	}

	// discover the in-cluster resources
	apiGroupResources, err := restmapper.GetAPIGroupResources(discoveryClient)
	if err != nil {
		logger.Error().Msgf("error retrieving api resources")
		panic(err)
	}

	// create a rest mapper that is used to convert GVK to GVR
	restMapper := restmapper.NewDiscoveryRESTMapper(apiGroupResources)

	// create the resources in deployment groups
	deploy := [][]*unstructured.Unstructured{
		resources.CustomResourceDefinitions(),
		resources.Namespaces(),
		resources.RBAC(),
		resources.Deployments(),
		resources.Services(),
		resources.OperatorDeploymentConfigs(),
		resources.Configs(),
	}

	// deploy the resources
	for _, group := range deploy {
		// create the resources
		for _, resource := range group {
			gvk := resource.GroupVersionKind()

			mapping, err := restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
			if err != nil {
				logger.Error().Msgf("error retrieving group version resource - %s", messageString(resource))
				panic(err)
			}
			gvr := mapping.Resource

			logger.Info().Msgf("creating resource - %s", messageString(resource))

			_, err = kubeClient.Resource(gvr).Namespace(resource.GetNamespace()).Apply(context.TODO(), resource.GetName(), resource, metav1.ApplyOptions{FieldManager: "platform-bootstrapper"})
			if err != nil {
				logger.Error().Msgf("error creating resource - %s", messageString(resource))
				panic(err)
			}

			// if we have extended the kubernetes api with a custom resource definition, we need to rediscover
			// the api resources
			if gvk.Kind == "CustomResourceDefinition" {
				// discover the in-cluster resources
				apiGroupResources, err := restmapper.GetAPIGroupResources(discoveryClient)
				if err != nil {
					logger.Error().Msgf("error retrieving api resources")
					panic(err)
				}

				// create a rest mapper that is used to convert GVK to GVR
				restMapper = restmapper.NewDiscoveryRESTMapper(apiGroupResources)
			}

			logger.Info().Msgf("waiting for resource to be ready - %s", messageString(resource))

			// wait for the resource to be ready
			if err := waitFor(resource, kubeClient, restMapper); err != nil {
				panic(err)
			}
		}
	}

	logger.Info().Msgf("successfully completed deployment")
}

func messageString(resource *unstructured.Unstructured) string {
	return fmt.Sprintf("group=%s, version=%s, kind=%s, namespace=%s, name=%s",
		resource.GetObjectKind().GroupVersionKind().Group,
		resource.GetObjectKind().GroupVersionKind().Version,
		resource.GetObjectKind().GroupVersionKind().Kind,
		resource.GetNamespace(),
		resource.GetName(),
	)
}

func waitFor(resource *unstructured.Unstructured, kubeClient *dynamic.DynamicClient, restMapper meta.RESTMapper) error {
	timeout, interval := time.After(waitTimeout), time.Tick(waitInterval)

	gvk := resource.GroupVersionKind()

	mapping, err := restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return fmt.Errorf("error retrieving group version resource - %s", messageString(resource))
	}
	gvr := mapping.Resource

	for {
		select {
		case <-timeout:
			return fmt.Errorf("timed out waiting for resource")
		case <-interval:
			resourceFromCluster, err := kubeClient.Resource(gvr).Namespace(resource.GetNamespace()).Get(
				context.TODO(),
				resource.GetName(),
				metav1.GetOptions{},
			)
			if err != nil {
				return fmt.Errorf("error getting resource - %s", messageString(resource))
			}

			ready, err := resourceutils.IsReady(resourceFromCluster)
			if err != nil {
				return fmt.Errorf("error waiting for resource to be ready, %w", err)
			}

			if ready {
				return nil
			}
		}
	}
}
