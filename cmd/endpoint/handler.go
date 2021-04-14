package endpoint

import (
	"cloud-final-project/cmd/restaurant"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func subscribe(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	restaurant.Subscribe(conn)
}

func register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newRestaurant restaurant.Restaurant
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newRestaurant)
	if err != nil {
		log.Println(err)
		return
	}

	restaurant.Register(newRestaurant)
}

func checkIn(w http.ResponseWriter, r *http.Request) {
	restaurant.CheckIn(parseRestaurantID(r))
}

func checkOut(w http.ResponseWriter, r *http.Request) {
	restaurant.CheckOut(parseRestaurantID(r))
}

func getRestaurantInfo(w http.ResponseWriter, r *http.Request) {
	info, err := json.Marshal(restaurant.GetInfo(parseRestaurantID(r)))
	if err != nil {
		w.WriteHeader(500)
	}

	fmt.Fprint(w, string(info))
}
