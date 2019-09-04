package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/roeyaus/drtest/db"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

func GetCabRidesForMedallions(medallions []string) ([]*db.CabRide, []string, error) {
	var cabRides []*db.CabRide
	var missingMedallions []string
	for _, m := range medallions {
		var cabRide *db.CabRide
		val2, err := client.Get(m).Result()
		if err == redis.Nil {
			fmt.Printf("%v does not exist\n", m)
			missingMedallions = append(missingMedallions, m)
		} else if err != nil {
			missingMedallions = append(missingMedallions, m)
		} else {
			if err = json.Unmarshal([]byte(val2), &cabRide); err != nil {
				missingMedallions = append(missingMedallions, m)
			}
			cabRides = append(cabRides, cabRide)
		}
	}

	return cabRides, missingMedallions, nil
}

func SetCabRides(cabRide []*db.CabRide) error {
	for _, cr := range cabRide {
		json, err := json.Marshal(cr)
		if err != nil {
			fmt.Printf("SetCabRides::Marshal failed %v", err)
		}
		err = client.Set(cr.Medallion, json, 30*time.Second).Err()
		if err != nil {
			fmt.Printf("SetCabRides::Set failed %v", err)
		}
	}

	return nil
}

func ClearCacheForMedallions(medallions []string) error {
	res := client.Del(medallions...)
	return res.Err()
}
