package cache

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/roeyaus/drtest/db"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

func GetCabRidesForMedallions(medallions []string) (*[]db.CabRide, error) {
	var cabRides *[]db.CabRide
	for _, m := range medallions {
		var cabRide *db.CabRide
		val2, err := client.Get(medallion).Result()
		if err == redis.Nil {
			fmt.Printf("%v does not exist\n", medallion)
			return nil, nil
		} else if err != nil {
			return nil, errors.Wrap(err, "GetCabRideForMedallion::Get failed")
		} else {
			if err = json.Unmarshal([]byte(val2), &cabRide); err != nil {
				return nil, errors.Wrap(err, "GetCabRideForMedallion::Unmarshal failed")
			}
			cabRides = append(cabRides, cabRide)
		}
	}
	
	return nil, nil
}

func SetCabRide(cabRide *db.CabRide) error {
	err := client.Set(cabRide.Medallion, cabRide, 0).Err()
	if err != nil {
		return errors.Wrap(err, "SetCabRide::Get failed")
	}
	return nil
}
