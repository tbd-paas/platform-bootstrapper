package resources

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Namespaces returns the namespace where the platform controllers will be deployed.
func Namespaces() []*unstructured.Unstructured {
	return []*unstructured.Unstructured{}
}

// CustomResourceDefinintions returns any custom resource definitions that the controllers
// must watch prior to starting.
func CustomResourceDefinitions() []*unstructured.Unstructured {
	return []*unstructured.Unstructured{
		customResourceDefinitionPlatformoperatorsDeployPlatformTbdIo,
		customResourceDefinitionPlatformconfigsDeployPlatformTbdIo,
	}
}

// RBAC returns any role-based access that the controllers need.
func RBAC() []*unstructured.Unstructured {
	return []*unstructured.Unstructured{
		serviceAccountPlatformConfigOperatorControllerManager,
		rolePlatformConfigOperatorLeaderElectionRole,
		roleBindingPlatformConfigOperatorLeaderElectionRolebinding,
		clusterRolePlatformConfigOperatorMetricsReader,
		clusterRolePlatformConfigOperatorProxyRole,
		clusterRolePlatformConfigOperatorManagerRole,
		clusterRoleBindingPlatformConfigOperatorProxyRolebinding,
		clusterRoleBindingPlatformConfigOperatorManagerRolebinding,
	}
}

// Deployments returns the actual deployment resource that runs the controller.
func Deployments() []*unstructured.Unstructured {
	return []*unstructured.Unstructured{deploymentPlatformConfigOperatorControllerManager}
}

// Services returns the actual service resource that runs the controller.
func Services() []*unstructured.Unstructured {
	return []*unstructured.Unstructured{servicePlatformConfigOperatorControllerManagerMetricsService}
}

// OperatorDeploymentConfigs returns the resource to control the platform operator deployments.
func OperatorDeploymentConfigs() []*unstructured.Unstructured {
	deployment := platformOperatorsConfig

	deployment.SetName("config")

	return []*unstructured.Unstructured{deployment}
}

// Configs returns the resource to control the platform configuration.
func Configs() []*unstructured.Unstructured {
	config := platformConfigConfig

	config.SetName("config")

	return []*unstructured.Unstructured{config}
}

// ApplyOrder returns the order in which resources need to be applied.
func ApplyOrder() []*unstructured.Unstructured {
	objects := []*unstructured.Unstructured{}

	for _, group := range [][]*unstructured.Unstructured{
		CustomResourceDefinitions(),
		Namespaces(),
		RBAC(),
		Deployments(),
		Services(),
		OperatorDeploymentConfigs(),
		Configs(),
	} {
		objects = append(objects, group...)
	}

	return objects
}

// DestroyOrder returns the order in which resources need to be destroyed.
func DestroyOrder() []*unstructured.Unstructured {
	forward := ApplyOrder()

	reverse := make([]*unstructured.Unstructured, len(forward))

	for i, j := len(forward)-1, 0; j < len(forward); i, j = i-1, j+1 {
		reverse[j] = forward[i]
	}

	return reverse
}
