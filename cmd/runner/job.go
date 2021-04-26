package runner

import (
	"cloud-final-project/cmd/restaurant"
	"cloud-final-project/cmd/timestream"
	"log"
	"time"
)

func SaveSnapshot() {
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			snapshot := restaurant.GetSnapshot()
			timestream.WriteSnapshot(snapshot)
			log.Println("save snapshot successfully")
		}
	}
}
