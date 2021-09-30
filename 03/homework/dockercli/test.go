package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func test() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalln(err)
	}

	stop := flag.NewFlagSet("stop", flag.ExitOnError)
	start := flag.NewFlagSet("start", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("stop or start subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "start":
		start.Parse(os.Args[2:])
	case "stop":
		stop.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	startId := strings.Join(start.Args(), "")
	stopID := strings.Join(stop.Args(), "")

	if stop.Parsed() {
		// Required Flags
		if stopID == "" {
			//stop.PrintDefaults()
			fmt.Printf("Container ID or Container Name must be defined: %s", stopID)
			os.Exit(1)
		}
		if err := cli.ContainerStop(ctx, stopID, nil); err != nil {
			log.Fatalln("Not able to stop container with err ", err.Error())
		}
		fmt.Printf("Stop container %s succeed", stopID)
	}

	if start.Parsed() {
		// Required Flags
		if startId == "" {
			//stop.PrintDefaults()
			fmt.Printf("Container ID or Container Name must be defined: %s", startId)
			os.Exit(1)
		}
		if err := cli.ContainerStart(ctx, startId, types.ContainerStartOptions{}); err != nil {
			log.Fatalln("Not able to start container with err ", err.Error())
		}
		fmt.Printf("Start container %s succeed", stopID)
	}

}
