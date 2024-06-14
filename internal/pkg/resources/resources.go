package resources

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

// Namespaces returns the namespace where the platform controllers will be deployed.
func Namespaces() []*unstructured.Unstructured {
	return []*unstructured.Unstructured{namespaceTbdOperatorsSystem}
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
