package endpoint

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

// upgrader for websocket connection
var upgrader websocket.Upgrader

var port string

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	port = fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
}

func ListenAndServe() {
	setHandler()

	log.Printf("Application is listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func setHandler() {
	fileServer := http.FileServer(http.Dir("./static"))

	// dashboard frontend
	http.Handle("/", fileServer)

	// register
	http.HandleFunc("/register", register)

	// checkin - checkout endpoint
	http.HandleFunc("/checkin", checkIn)
	http.HandleFunc("/checkout", checkOut)

	// restaurant info
	http.HandleFunc("/info", getRestaurantInfo)

	// subscribe value change
	http.HandleFunc("/subscribe", subscribe)

}
