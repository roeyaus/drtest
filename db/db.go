package db

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var db *sqlx.DB

type CabRide struct {
	Medallion  string    `db:"medallion"`
	PickupDate time.Time `db:"pickup_date"`
	NumTrips   int       `db:"num_trips"`
}

func init() {
	var err error
	db, err = sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/mydb?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
}

func GetCabRidesForMedallions(medallions []string) ([]*CabRide, error) {
	cabRides := []*CabRide{}
	fmt.Printf("%+v", medallions)
	query, args, err := sqlx.In("SELECT medallion, DATE(pickup_datetime) as pickup_date, COUNT(medallion) as num_trips FROM cab_trip_data WHERE medallion IN (?) GROUP BY medallion, DATE(pickup_datetime)", medallions)
	if err != nil {
		return nil, errors.Wrap(err, "GetCabRidesForMedallions::In failed")
	}

	if err = db.Select(&cabRides, db.Rebind(query), args...); err != nil {
		return nil, errors.Wrap(err, "GetCabRidesForMedallions::Select failed")
	}
	return cabRides, nil
}
