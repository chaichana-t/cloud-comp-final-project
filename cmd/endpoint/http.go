package endpoint

import (
	syncService "cloud-final-project/cmd/sync-service"
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

	http.HandleFunc("/checkin", checkIn)
	http.HandleFunc("/checkout", checkOut)

	go listen(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")))
}

func getRestaurantID(r *http.Request) string {
	return r.URL.Query()["rid"][0]
}

func listen(port string) {
	log.Printf("Application is listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func checkIn(w http.ResponseWriter, r *http.Request) {
	syncService.Increase(getRestaurantID(r))
}

func checkOut(w http.ResponseWriter, r *http.Request) {
	syncService.Decrease(getRestaurantID(r))
}
