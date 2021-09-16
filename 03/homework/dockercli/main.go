package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	term "github.com/nsf/termbox-go"
	"log"
	"os"
)

func main() {

	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalln(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Enter F5 to toggle between options, press ESC or CTRL+C button to quit")

	for i := 1; i <= 3; i++ {
		ev := term.PollEvent()
		if ev.Key == term.KeyF5 && i == 1 {
			for _, container := range containers {
				if container.State == "running" {
					fmt.Println(container.ID, container.Names, container.Image, container.Ports, container.State)
				}
			}
			fmt.Println("--------------------------")
		}

		if ev.Key == term.KeyF5 && i == 2 {
			for _, container := range containers {
				if container.State == "exited" {
					fmt.Println(container.ID, container.Names, container.Image, container.Ports, container.State)
				}
			}
			fmt.Println("--------------------------")
		}

		if ev.Key == term.KeyF5 && i == 3 {
			i = 0
			for _, container := range containers {
				fmt.Println(container.ID, container.Names, container.Image, container.Ports, container.State)
			}
			fmt.Println("--------------------------")
		}

		if ev.Key == term.KeyEsc || ev.Key == term.KeyCtrlC {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}
}
