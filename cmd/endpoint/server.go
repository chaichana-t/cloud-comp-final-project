package endpoint

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/checkin", checkIn)
	http.HandleFunc("/checkout", checkOut)

	http.HandleFunc("/info", getRestaurantInfo)

	http.HandleFunc("/subscribe", subscribe)

	go listen(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")))
}

