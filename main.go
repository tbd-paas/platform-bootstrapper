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

	// apply or destroy the objects based on the action
	if os.Getenv("BOOTSTRAP_ACTION") == "destroy" {
		if err := b.RunAction(b.Destroy, resources.DestroyOrder()...); err != nil {
			panic(fmt.Errorf("error destroying bootstrapped resources - %w", err))
		}
	} else {
		if err := b.RunAction(b.Apply, resources.ApplyOrder()...); err != nil {
			panic(fmt.Errorf("error applying bootstrapped resources - %w", err))
		}
	}

	b.Log.Info().Msg("successfully completed bootstrapping")
}
