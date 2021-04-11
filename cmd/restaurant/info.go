package restaurant

type Restaurant struct {
	Name string
	MaxCustomer int
}

var mockRestaurants = map[string]*Restaurant{
	"cad25750-9b14-11eb-a8b3-0242ac130003": {"Restaurant A", 20},
	"d936f9f4-9b14-11eb-a8b3-0242ac130003": {"Restaurant B", 50},
	"de42536c-9b14-11eb-a8b3-0242ac130003": {"Restaurant C", 30},
}

func GetInfo(restaurantID string) *Restaurant {
	return mockRestaurants[restaurantID]
}