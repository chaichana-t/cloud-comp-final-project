package sync_service

import (
	"github.com/go-redis/redis/v7"
	"log"
	"strconv"
)

var cache map[string]int

func Subscribe(key string) chan int {
	pubsub := client.Subscribe(key)
	payloads := make(chan int)

	go subscribeLoop(pubsub, payloads)

	log.Printf("Subscribe \"%s\" successfully.\n", key)

	return payloads
}

func subscribeLoop(pubsub *redis.PubSub, payloads chan int) {
	for {
		message, err := pubsub.ReceiveMessage()
		if err != nil {
			return
		}

		value, _ := strconv.Atoi(message.Payload)
		if value != cache[message.Channel] {
			cache[message.Channel] = value
		}
		payloads <- value
	}
}

func Get(key string) int {
	return cache[key]
}