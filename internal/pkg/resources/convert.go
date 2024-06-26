package resources

import (
	"fmt"

	"github.com/nukleros/operator-builder-tools/pkg/resources"
	"github.com/tbd-paas/platform-config-operator/apis/deploy/v1alpha1"
	"github.com/tbd-paas/tbd-cli/sdk/api/interfaces"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// ConvertInputToPlatformConfig converts a user input configuration, from its interface, into
// a PlatformConfig object.
func ConvertInputToPlatformConfig(input interfaces.Config) (*unstructured.Unstructured, error) {
	platformConfig := &v1alpha1.PlatformConfig{}

	// convert the default generated unstructured platform config to its proper type
	if err := resources.ToTyped(platformConfig, platformConfigConfig); err != nil {
		return nil, fmt.Errorf("unable to convert user input config to platform config - %w", err)
	}

	// set fields from the user input
	setDeploymentSize(input, platformConfig)
	setLocal(input, platformConfig)

	// convert to unstructured and return
	return resources.ToUnstructured(platformConfig)
}

// setDeploymentSize sets the deployment size for a PlatformConfig.
func setDeploymentSize(source interfaces.Config, destination *v1alpha1.PlatformConfig) {
	destination.Spec.Platform.Certificates.DeploymentSize = source.GetDeploymentSize()
	destination.Spec.Platform.Identity.DeploymentSize = source.GetDeploymentSize()
}

// setLocal sets the local value for a config.
func setLocal(source interfaces.Config, destination *v1alpha1.PlatformConfig) {
	destination.Spec.Cloud.Local = source.GetLocal()
}
