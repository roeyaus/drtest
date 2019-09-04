package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCabRidesForMedallions(t *testing.T) {
	rides, err := GetCabRidesForMedallions([]string{})
	assert.Nil(t, err)
	assert.Equal(t, []*CabRide{}, rides, "rides should be empty")

	rides, err = GetCabRidesForMedallions([]string{"D7D598CD99978BD012A87A76A7C891B7"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(rides), "there should be 1 medallion")

	rides, err = GetCabRidesForMedallions([]string{"D7D598CD99978BD012A87A76A7C891B7"})
	assert.Nil(t, err)
	assert.Equal(t, 3, rides[0].NumTrips, "there should be 3 trips for this medallion")

}
