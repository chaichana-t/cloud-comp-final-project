package main

import (
	syncService "cloud-final-project/cmd/sync-service"
	"log"
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
