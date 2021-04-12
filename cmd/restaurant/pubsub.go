package restaurant

import (
	syncService "cloud-final-project/cmd/sync-service"
	"fmt"
	"strconv"

	"github.com/gorilla/websocket"
)

func CheckIn(restaurantID string) {
	numberOfCustomer, err := strconv.Atoi(syncService.Get(restaurantID))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if numberOfCustomer >= *&GetInfo(restaurantID).MaxCustomer {
		fmt.Println("Full capacity")
		return
	}

	value := syncService.Increase(restaurantID)
	syncService.Publish(redisPayloadChannel, constructPayload(restaurantID, value))
}

func CheckOut(restaurantID string) {
	numberOfCustomer, err := strconv.Atoi(syncService.Get(restaurantID))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if numberOfCustomer <= 0 {
		fmt.Println("Conflict")
		return
	}

	value := syncService.Decrease(restaurantID)
	syncService.Publish(redisPayloadChannel, constructPayload(restaurantID, value))
}

func Subscribe(connection *websocket.Conn) {
	clients[connection.RemoteAddr().String()] = connection

	for _, message := range messages {
		connection.WriteMessage(websocket.TextMessage, []byte(message))
	}
}

func listenRedisUpdate() {
	go syncService.Subscribe(redisPayloadChannel, payloadChannel)

	for {
		payloadHandler(<-payloadChannel)
	}
}
