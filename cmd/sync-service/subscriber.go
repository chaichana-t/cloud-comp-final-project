package sync_service

import (
	"log"
)

func Subscribe(pubsubChannel string, payloadChannel chan string) {
	pubsub := client.Subscribe(pubsubChannel)

	log.Printf("Subscribe \"%s\" channel successfully.\n", pubsubChannel)
	for {
		message, err := pubsub.ReceiveMessage()
		if err != nil {
			log.Print(err)
			continue
		}
		payloadChannel <- message.Payload
	}
}
