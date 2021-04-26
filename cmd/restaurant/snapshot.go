package restaurant

import syncService "cloud-final-project/cmd/sync-service"

func GetSnapshot() map[string]string {
	snapshot := map[string]string{}

	restaurantIDs := getRestaurants()
	for _, id := range restaurantIDs {
		snapshot[GetInfo(id).Name] = syncService.Get(id)
	}

	return snapshot
}
