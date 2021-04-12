package endpoint

import "net/http"

func parseRestaurantID(r *http.Request) string {
	return r.URL.Query()["rid"][0]
}
