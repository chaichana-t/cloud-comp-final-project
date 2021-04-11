package restaurant

import (
	syncService "cloud-final-project/cmd/sync-service"
	"github.com/gorilla/websocket"
	"log"
)

var messages map[string]string
var payloadChannel chan string

var clients map[string]*websocket.Conn

const redisPayloadChannel = "value-change"

func init() {
	messages = make(map[string]string)
	payloadChannel = make(chan string)

	restaurants := getRestaurants()

	for _, restaurantID := range restaurants {
		payloadHandler(constructPayload(restaurantID, syncService.Get(restaurantID)))
	}

	go listenRedisUpdate()
}

func getRestaurants() []string {
	return []string{"1234"}
}

func constructPayload(restaurantID string, value string) string {
	// restaurantID is UUID (length: 36)
	return restaurantID + value
}

func payloadHandler(payload string) {
	messages[payload[:36]] = payload
	broadcastMessage(payload)
}

func broadcastMessage(message string) {
	for addr, connection := range clients {
		if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println(err)
			delete(clients, addr)
		}
	}
}