package main

import (
	"github.com/kataras/iris"
	"github.com/roeyaus/drtest/db"
	"github.com/roeyaus/drtest/cache"
)

type Request struct {
	Medallions []string `json:"medallions"`
}

func main() {
	app := iris.Default()
	app.Post("/trips", func(ctx iris.Context) {
		var bodyJSON Request
		if err := ctx.ReadJSON(&bodyJSON); err != nil {
			ctx.StatusCode(500)
			ctx.Writef("error : %v", err)
			return
		}
		if len(bodyJSON.Medallions) == 0 {
			ctx.StatusCode(401)
			ctx.Writef("error : no medallions provided")
			return
		}
		ctx.Writef("provided %v medallions", len(bodyJSON.Medallions))
		//first check cache
		//cache.GetCabRideForMedallion()
		cabRides, err := db.GetCabRidesForMedallions(bodyJSON.Medallions)
		if err != nil {
			ctx.StatusCode(500)
			ctx.Writef("error : %v", err)
			return
		}
		ctx.JSON(cabRides)
	})

	app.Run(iris.Addr(":8080"))
}
