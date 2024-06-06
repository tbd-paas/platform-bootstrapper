package resources

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

// Namespace returns the namespace where the platform controllers will be deployed.
func Namespace() *unstructured.Unstructured {
	return namespaceTbdOperatorsSystem
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

// Deployment returns the actual deployment resource that runs the controller.
func Deployment() *unstructured.Unstructured {
	return deploymentPlatformConfigOperatorControllerManager
}

// Service returns the actual service resource that runs the controller.
func Service() *unstructured.Unstructured {
	return servicePlatformConfigOperatorControllerManagerMetricsService
}

// OperatorDeploymentConfig returns the resource to control the platform operator deployments.
func OperatorDeploymentConfig() *unstructured.Unstructured {
	deployment := platformOperatorsConfig

	deployment.SetName("config")

	return deployment
}

// Config returns the resource to control the platform configuration.
func Config() *unstructured.Unstructured {
	config := platformConfigConfig

	config.SetName("config")

	return config
}
