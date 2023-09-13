package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/client"

	"dps/internal/config"
	"dps/internal/docker"
)

func main() {
	var err error

	var ctx = context.Background()

	appConfig := config.Init(ctx)
	ctx = config.SetConfig(ctx, appConfig)

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: container, image, network, volume")
		return
	}

	// Get new docker client
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	command := os.Args[1]

	switch command {
	case "container":
		docker.ShowContainers(ctx, cli)
	case "image":
		docker.ShowImages(ctx, cli)
	case "network":
		docker.ShowNetworks(ctx, cli)
	case "volume":
		docker.ShowVolumes(ctx, cli)
	default:
		fmt.Printf("Invalid command: %s\n", command)
	}

}
