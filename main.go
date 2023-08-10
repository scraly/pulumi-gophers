package main

import (
	"fmt"

	"github.com/pulumi/pulumi-docker/sdk/v3/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		//Configuration
		cfg := config.New(ctx, "")
		gophersAPIPort := cfg.RequireFloat64("gophersAPIPort")
		gophersAPIWatcherPort := cfg.RequireFloat64("gophersAPIWatcherPort")
		//TODO: temp
		_ = gophersAPIPort + gophersAPIWatcherPort

		// Pull the Gophers API image
		gophersAPIImageName := "gophers-api"
		gophersAPIImage, err := docker.NewRemoteImage(ctx, fmt.Sprintf("%v-image", gophersAPIImageName), &docker.RemoteImageArgs{
			Name: pulumi.String("scraly/" + gophersAPIImageName + ":latest"),
		})
		if err != nil {
			return err
		}
		ctx.Export("gophersAPIDockerImage", gophersAPIImage.Name)

		// Pull the Gophers API Watcher (frontend/UI) image
		gophersAPIWatcherImageName := "gophers-api-watcher"
		gophersAPIWatcherImage, err := docker.NewRemoteImage(ctx, fmt.Sprintf("%v-image", gophersAPIWatcherImageName), &docker.RemoteImageArgs{
			Name: pulumi.String("scraly/" + gophersAPIWatcherImageName + ":latest"),
		})
		if err != nil {
			return err
		}
		ctx.Export("gophersAPIWatcherDockerImage", gophersAPIWatcherImage.Name)

		// Create a Docker network
		network, err := docker.NewNetwork(ctx, "network", &docker.NetworkArgs{
			Name: pulumi.String(fmt.Sprintf("services-%v", ctx.Stack())),
		})
		if err != nil {
			return err
		}
		ctx.Export("containerNetwork", network.Name)

		return nil
	})
}
