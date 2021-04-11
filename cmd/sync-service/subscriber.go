package sync_service

import (
	"log"
	"strconv"
)

func Subscribe(key string) chan int {
	pubsub := client.Subscribe(key)
	payloads := make(chan int)
	go func() {
		for {
			message, err := pubsub.ReceiveMessage()
			if err != nil {
				return
			}
			value, _ := strconv.Atoi(message.Payload)
			payloads <- value
		}
	}()
	log.Printf("Subscribe \"%s\" successfully.\n", key)
	return payloads
}

