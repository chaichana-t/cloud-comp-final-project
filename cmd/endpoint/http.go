package endpoint

import (
	"cloud-final-project/cmd/restaurant"
	"fmt"
	"log"
	"net/http"
)

func getRestaurantID(r *http.Request) string {
	return r.URL.Query()["rid"][0]
}

func listen(port string) {
	log.Printf("Application is listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func checkIn(w http.ResponseWriter, r *http.Request) {
	restaurant.CheckIn(getRestaurantID(r))
}

func checkOut(w http.ResponseWriter, r *http.Request) {
	restaurant.CheckOut(getRestaurantID(r))
}

func getRestaurant(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, restaurant.GetName(getRestaurantID(r)))
}