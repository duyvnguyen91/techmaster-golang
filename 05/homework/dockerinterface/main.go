package main

import (
	"log"
	"techmaster-golang/05/homework/dockerinterface/container"
)

func main() {
	container.ActiontListContainer()
	if err := container.ActionStopContainer("mysql"); err != nil {
		log.Fatalf("Error Stoping Docker")
	}
	if err := container.ActionStartContainer("blissful_taussig"); err != nil {
		log.Fatalf("Error Starting Docker")
	}

}
