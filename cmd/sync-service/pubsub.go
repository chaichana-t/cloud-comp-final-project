package sync_service

import "log"

func Increase(key string) string {
	return string(client.Incr(key).Val())
}

func Decrease(key string) string {
	return string(client.Decr(key).Val())
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
	return result.Val()
}