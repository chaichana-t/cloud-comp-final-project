package endpoint

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/checkIn", checkIn)
	http.HandleFunc("/checkOut", checkOut)

	go listen(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")))
}

func listen(port string) {
	log.Printf("Application is listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func checkIn(w http.ResponseWriter, r *http.Request) {

}

func checkOut(w http.ResponseWriter, r *http.Request) {

}
