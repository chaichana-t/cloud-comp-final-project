package main

import (
	syncService "cloud-final-project/cmd/sync-service"
	"log"

	_ "cloud-final-project/cmd/endpoint"
)

func main() {
	c := syncService.Subscribe("A")

	syncService.Increase("A")
	syncService.Increase("A")
	syncService.Increase("B")
	syncService.Increase("B")

	for {
		log.Println(<-c)
	}
}
