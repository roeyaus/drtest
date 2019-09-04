package cache

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/roeyaus/drtest/db"
	"github.com/stretchr/testify/assert"
)

func TestGetCabRidesForMedallions(t *testing.T) {
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	crTest := db.CabRide{
		Medallion:  "medallion1",
		PickupDate: then,
		NumTrips:   8,
	}

	json, err := json.Marshal(crTest)
	err = client.Set("medallion1", json, 0).Err()
	assert.Nil(t, err)
	rides, notCached, err := GetCabRidesForMedallions([]string{"medallion1"})
	assert.Equal(t, &crTest, rides[0], "this isn't what we stored")
	assert.Equal(t, []string{}, notCached, "notCached should be empty")

	rides, notCached, err = GetCabRidesForMedallions([]string{"medallion1", "medallion2"})
	assert.Equal(t, &crTest, rides[0], "this isn't what we stored")
	assert.Equal(t, []string{"medallion2"}, notCached, "notCached should equal medallion2")
}
