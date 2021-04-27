package restaurant

import (
	syncService "cloud-final-project/cmd/sync-service"
	"fmt"
	"strconv"

	"github.com/gorilla/websocket"
)

func Register(restaurant Restaurant) {
	mockRestaurants[restaurant.Name] = &restaurant
}

func CheckIn(restaurantID string) bool {
	numberOfCustomer, err := strconv.Atoi(syncService.Get(restaurantID))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if numberOfCustomer >= *&GetInfo(restaurantID).MaxCustomer {
		fmt.Println("Full capacity")
		return false
	}

	value := syncService.Increase(restaurantID)
	syncService.Publish(redisPayloadChannel, constructPayload(restaurantID, value))
	return true
}

func CheckOut(restaurantID string) bool {
	numberOfCustomer, err := strconv.Atoi(syncService.Get(restaurantID))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if numberOfCustomer <= 0 {
		fmt.Println("Conflict")
		return false
	}

	value := syncService.Decrease(restaurantID)
	syncService.Publish(redisPayloadChannel, constructPayload(restaurantID, value))
	return true
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
