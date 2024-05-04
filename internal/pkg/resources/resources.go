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
		serviceAccountPlatformConfigOperatorcontrollerManager,
		rolePlatformConfigOperatorleaderElectionRole,
		roleBindingPlatformConfigOperatorleaderElectionRolebinding,
		clusterRolePlatformConfigOperatormetricsReader,
		clusterRolePlatformConfigOperatorproxyRole,
		clusterRolePlatformConfigOperatormanagerRole,
		clusterRoleBindingPlatformConfigOperatorproxyRolebinding,
		clusterRoleBindingPlatformConfigOperatormanagerRolebinding,
	}
}

// Deployment returns the actual deployment resource that runs the controller.
func Deployment() *unstructured.Unstructured {
	return deploymentPlatformConfigOperatorcontrollerManager
}

// Service returns the actual service resource that runs the controller.
func Service() *unstructured.Unstructured {
	return servicePlatformConfigOperatorcontrollerManagerMetricsService
}

// OperatorDeploymentConfig returns the resource to control the platform operator deployments.
func OperatorDeploymentConfig() *unstructured.Unstructured {
	deployment := platformOperatorsPlatformoperatorsSample

	deployment.SetName("config")

	return deployment
}
