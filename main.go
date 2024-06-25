package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tbd-paas/platform-bootstrapper/internal/pkg/bootstrapper"
	"github.com/tbd-paas/platform-bootstrapper/internal/pkg/resources"
)

func main() {
	// create the client
	client, err := bootstrapper.NewClient(context.Background(), os.Getenv("KUBECONFIG"))
	if err != nil {
		panic(fmt.Errorf("error creating client - %w", err))
	}

	// create the bootstrapper
	options := []bootstrapper.Option{}
	if os.Getenv("DEBUG") == "true" {
		options = append(options, bootstrapper.WithDebug)
	}

	b := bootstrapper.NewBootstrapper(client, options...)

	action := os.Getenv("BOOTSTRAP_ACTION")
	switch action {
	case "destroy":
		if err := b.RunAction(b.Destroy, resources.DestroyOrder()...); err != nil {
			panic(fmt.Errorf("error destroying bootstrapped resources - %w", err))
		}
	case "apply":
		if err := b.RunAction(b.Apply, resources.ApplyOrder()...); err != nil {
			panic(fmt.Errorf("error applying bootstrapped resources - %w", err))
		}
	default:
		panic(fmt.Errorf("unknown action [%s], only 'apply' and 'destroy' are supported", action))
	}

	b.Log.Info().Msgf("successfully completed bootstrapping for action: [%s]", action)
}
