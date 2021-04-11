package restaurant

import (
	syncService "cloud-final-project/cmd/sync-service"
	"github.com/gorilla/websocket"
)

func CheckIn(restaurantID string) {
	value := syncService.Increase(restaurantID)
	syncService.Publish(redisPayloadChannel, constructPayload(restaurantID, value))
}

func CheckOut(restaurantID string) {
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
