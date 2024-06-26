package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/tbd-paas/tbd-cli/sdk/api"
	"github.com/tbd-paas/tbd-cli/sdk/api/interfaces"

	"github.com/tbd-paas/platform-bootstrapper/internal/pkg/bootstrapper"
	"github.com/tbd-paas/platform-bootstrapper/internal/pkg/resources"
)

const (
	envKubeconfigFile         = "KUBECONFIG"
	envDebug                  = "BOOTSTRAP_DEBUG"
	envBootstrapAction        = "BOOTSTRAP_ACTION"
	envBootstrapConfigFile    = "BOOTSTRAP_CONFIG_FILE"
	envBootstrapConfigVersion = "BOOTSTRAP_CONFIG_VERSION"
)

const (
	defaultConfigFileLocation = "/opt/platform-bootstrapper-config/config.yaml"
	defaultConfigFileVersion  = api.VersionV0
)

func main() {
	// create the client
	client, err := bootstrapper.NewClient(context.Background(), os.Getenv(envKubeconfigFile))
	if err != nil {
		panic(fmt.Errorf("error creating client - %w", err))
	}

	// create the bootstrapper
	options := []bootstrapper.Option{}
	if os.Getenv(envDebug) == "true" {
		options = append(options, bootstrapper.WithDebug)
	}

	b := bootstrapper.NewBootstrapper(client, options...)

	// if we have a supplied user-input configuration, use it otherwise grab
	// the default platformconfig resource
	configVersion := os.Getenv(envBootstrapConfigVersion)
	if configVersion == "" {
		configVersion = string(defaultConfigFileVersion)
	}

	configFile := os.Getenv(envBootstrapConfigFile)
	if configFile == "" {
		configFile = defaultConfigFileLocation
	}

	var input interfaces.Config
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		input = api.NewConfig(api.Version(configVersion))
	} else {
		input, err = api.NewConfigFromFile(api.Version(configVersion), configFile)
		if err != nil {
			panic(fmt.Errorf("error creating user input config from file [%s] - %w", configFile, err))
		}
	}

	platformConfig, err := resources.ConvertInputToPlatformConfig(input)
	if err != nil {
		panic(fmt.Errorf("error converting user input to platform config - %w", err))
	}

	// run the bootstrapper action
	action := os.Getenv(envBootstrapAction)
	switch action {
	case "destroy":
		if err := b.RunAction(b.Destroy, resources.DestroyOrder(platformConfig)...); err != nil {
			panic(fmt.Errorf("error destroying bootstrapped resources - %w", err))
		}
	case "apply":
		if err := b.RunAction(b.Apply, resources.ApplyOrder(platformConfig)...); err != nil {
			panic(fmt.Errorf("error applying bootstrapped resources - %w", err))
		}
	default:
		panic(fmt.Errorf("unknown action [%s], only 'apply' and 'destroy' are supported", action))
	}

	b.Log.Info().Msgf("successfully completed bootstrapping for action: [%s]", action)
}
