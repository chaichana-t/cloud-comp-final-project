package sync_service

import (
	"log"
	"strconv"
)

func Increase(key string, max int) (string, bool) {
	result := int(client.Incr(key).Val())
	if result > max {
		client.Decr(key)
		return strconv.Itoa(result-1), false
	}
	return strconv.Itoa(result), true
}

func Decrease(key string) string {
	result := int(client.Decr(key).Val())
	if result < 0 {
		client.Incr(key)
		return strconv.Itoa(result+1)
	}
	return strconv.Itoa(result)
}

func Publish(pubsubChannel string, message string) {
	client.Publish(pubsubChannel, message)
}

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

func Get(key string) string {
	result := client.Get(key)
	if result.Val() == "" {
		return "0"
	}
	return result.Val()
}
