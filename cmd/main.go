package main

import (
	syncService "cloud-final-project/cmd/sync-service"
	"log"

	_ "cloud-final-project/cmd/endpoint"
)

func main() {
	c := syncService.Subscribe("1234")

	for {
		log.Println(<-c)
	}
}
