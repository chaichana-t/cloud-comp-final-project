package endpoint

import (
	"cloud-final-project/cmd/restaurant"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader websocket.Upgrader

func init() {
	upgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
}

func subscribe(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	restaurant.Subscribe(conn)
}
