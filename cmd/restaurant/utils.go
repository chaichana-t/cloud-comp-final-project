package restaurant

import (
	"github.com/gorilla/websocket"
	"log"
)

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


