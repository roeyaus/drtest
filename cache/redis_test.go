package cache

import (
	"testing"
	"github.com/roeyaus/drtest/db"
	"github.com/stretchr/testify/assert"
)

func TestGetCabRidesForMedallions() {

	err := client.Set("medallion1", db.CabRide{
		Medallion:"medallion1",
		PickupDate: time.Now(),
		NumTrips:8,
	}, 0).Err()

}