package restaurant

type Restaurant struct {
	Name        string `json:"name"`
	MaxCustomer int    `json:"maxCustomer"`
}

var mockRestaurants = map[string]*Restaurant{
	"cad25750-9b14-11eb-a8b3-0242ac130003": {"Starbucks (La Villa Ari)", 62},
	"d936f9f4-9b14-11eb-a8b3-0242ac130003": {"Starbucks (Siam Square One)", 309},
	"de42536c-9b14-11eb-a8b3-0242ac130003": {"Bonchon (Siam Center)", 83},
}

func GetInfo(restaurantID string) *Restaurant {
	return mockRestaurants[restaurantID]
}

func getRestaurants() []string {
	var restaurantIDs []string

	for restaurantID, _ := range mockRestaurants {
		restaurantIDs = append(restaurantIDs, restaurantID)
	}

	return restaurantIDs
}
