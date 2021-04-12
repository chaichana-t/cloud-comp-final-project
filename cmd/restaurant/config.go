package restaurant

import (
	syncService "cloud-final-project/cmd/sync-service"
	"github.com/gorilla/websocket"
)

var messages map[string]string
var payloadChannel chan string

var clients map[string]*websocket.Conn

const redisPayloadChannel = "value-change"

func init() {
	messages = make(map[string]string)
	payloadChannel = make(chan string)
	clients = make(map[string]*websocket.Conn)

	restaurants := getRestaurants()

	for _, restaurantID := range restaurants {
		payloadHandler(constructPayload(restaurantID, syncService.Get(restaurantID)))
	}

	go listenRedisUpdate()
}
