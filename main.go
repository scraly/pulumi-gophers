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
		protocol := "http://"
		//tag := "latest"
		//for GitPod!
		tag := "linux-amd64"

		cfg := config.New(ctx, "")
		gophersAPIPort := cfg.RequireFloat64("gophersAPIPort")
		gophersAPIWatcherPort := cfg.RequireFloat64("gophersAPIWatcherPort")
		//TODO: temp
		_ = gophersAPIWatcherPort

		// Pull the Gophers API image
		gophersAPIImageName := "gophers-api"
		//TODO: set platform?
		gophersAPIImage, err := docker.NewRemoteImage(ctx, fmt.Sprintf("%v-image", gophersAPIImageName), &docker.RemoteImageArgs{
			Name: pulumi.String("scraly/" + gophersAPIImageName + ":" + tag),
		})
		if err != nil {
			return err
		}
		ctx.Export("gophersAPIDockerImage", gophersAPIImage.Name)

		// Pull the Gophers API Watcher (frontend/UI) image
		gophersAPIWatcherImageName := "gophers-api-watcher"
		gophersAPIWatcherImage, err := docker.NewRemoteImage(ctx, fmt.Sprintf("%v-image", gophersAPIWatcherImageName), &docker.RemoteImageArgs{
			Name: pulumi.String("scraly/" + gophersAPIWatcherImageName + ":" + tag),
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

		// Create the gophers API container
		// Use _ instead of a variable name since this container isn't referenced

		//TODO: set platform
		_, err = docker.NewContainer(ctx, "gophers-api", &docker.ContainerArgs{
			Name:  pulumi.String(fmt.Sprintf("gophers-api-%v", ctx.Stack())),
			Image: gophersAPIImage.RepoDigest,
			Ports: &docker.ContainerPortArray{
				&docker.ContainerPortArgs{
					Internal: pulumi.Int(gophersAPIPort),
					External: pulumi.Int(gophersAPIPort),
				},
			},
			NetworksAdvanced: &docker.ContainerNetworksAdvancedArray{
				&docker.ContainerNetworksAdvancedArgs{
					Name: network.Name,
					Aliases: pulumi.StringArray{
						pulumi.String(fmt.Sprintf("gophers-api-%v", ctx.Stack())),
					},
				},
			},
		})
		if err != nil {
			return err
		}

		// Create the frontend container
		_, err = docker.NewContainer(ctx, "gophers-api-watcher", &docker.ContainerArgs{
			Name:  pulumi.String(fmt.Sprintf("gophers-api-watcher-%v", ctx.Stack())),
			Image: gophersAPIWatcherImage.RepoDigest,
			Ports: &docker.ContainerPortArray{
				&docker.ContainerPortArgs{
					Internal: pulumi.Int(gophersAPIWatcherPort),
					External: pulumi.Int(gophersAPIWatcherPort),
				},
			},
			Envs: pulumi.StringArray{
				pulumi.String(fmt.Sprintf("PORT=%v", gophersAPIWatcherPort)),
				pulumi.String(fmt.Sprintf("HTTP_PROXY=backend-%v:%v", ctx.Stack(), gophersAPIPort)),
				pulumi.String(fmt.Sprintf("PROXY_PROTOCOL=%v", protocol)),
			},
			NetworksAdvanced: &docker.ContainerNetworksAdvancedArray{
				&docker.ContainerNetworksAdvancedArgs{
					Name: network.Name,
					Aliases: pulumi.StringArray{
						pulumi.String(fmt.Sprintf("gophers-api-watcher-%v", ctx.Stack())),
					},
				},
			},
		})
		if err != nil {
			return err
		}

		return nil
	})
}
