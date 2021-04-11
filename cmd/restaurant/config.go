package restaurant

import (
	syncService "cloud-final-project/cmd/sync-service"
	"github.com/gorilla/websocket"
	"log"
)

var messages map[string]string
var payloadChannel chan string

var clients map[string]*websocket.Conn

var mockRestaurantName map[string]string

const redisPayloadChannel = "value-change"

func init() {
	messages = make(map[string]string)
	payloadChannel = make(chan string)
	clients = make(map[string]*websocket.Conn)

	mockRestaurantName = map[string]string{
		"cad25750-9b14-11eb-a8b3-0242ac130003": "Restaurant A",
		"d936f9f4-9b14-11eb-a8b3-0242ac130003": "Restaurant B",
		"de42536c-9b14-11eb-a8b3-0242ac130003": "Restaurant C",
	}

	restaurants := getRestaurants()

	for _, restaurantID := range restaurants {
		payloadHandler(constructPayload(restaurantID, syncService.Get(restaurantID)))
	}

	go listenRedisUpdate()
}

func getRestaurants() []string {
	return []string{"cad25750-9b14-11eb-a8b3-0242ac130003", "d936f9f4-9b14-11eb-a8b3-0242ac130003", "de42536c-9b14-11eb-a8b3-0242ac130003"}
}

func GetName(restaurantID string) string {
	return mockRestaurantName[restaurantID]
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