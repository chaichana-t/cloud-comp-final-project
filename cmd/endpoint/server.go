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

	http.HandleFunc("/checkin", checkIn)
	http.HandleFunc("/checkout", checkOut)
	http.HandleFunc("/subscribe", subscribe)

	go listen(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")))
}

