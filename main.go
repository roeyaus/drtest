package main

import (
	"io/ioutil"
	"strings"

	"github.com/kataras/iris"
	"github.com/roeyaus/drtest/cache"
	"github.com/roeyaus/drtest/db"
)

type Request struct {
	Medallions []string `json:"medallions"`
}

func main() {
	app := iris.Default()
	app.Post("/trips", func(ctx iris.Context) {
		var err error
		//check clearcache param
		noCache, _ := ctx.URLParamBool("nocache")
		rawData, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			ctx.StatusCode(500)
			ctx.Writef("error : %v\n", err)
			return
		}
		medallions := strings.Split(string(rawData), ",")
		if len(medallions) == 0 {
			ctx.StatusCode(401)
			ctx.Writef("error : no medallions provided\n")
			return
		}

		var cabRides []*db.CabRide
		var notCached []string
		//should we clear the cache for these medallions first?
		if !noCache {
			//first check cache
			cabRides, notCached, err = cache.GetCabRidesForMedallions(medallions)
			if err != nil {
				ctx.StatusCode(500)
				ctx.Writef("error : %v\n", err)
				return
			}
		} else {
			notCached = medallions
		}

		cabRidesFromDB, err := db.GetCabRidesForMedallions(notCached)
		if err != nil {
			ctx.StatusCode(500)
			ctx.Writef("error : %v\n", err)
			return
		}
		cabRides = append(cabRides, cabRidesFromDB...)
		_ = cache.SetCabRides(cabRides)
		rides := map[string]*db.CabRide{}
		for _, r := range cabRides {
			rides[r.Medallion] = r
		}
		ctx.JSON(rides)
	})

	app.Post("/clearcache", func(ctx iris.Context) {
		rawData, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			ctx.StatusCode(500)
			ctx.Writef("error : %v\n", err)
			return
		}
		medallions := strings.Split(string(rawData), ",")
		if len(medallions) == 0 {
			ctx.StatusCode(401)
			ctx.Writef("error : no medallions provided\n")
			return
		}
		if err := cache.ClearCacheForMedallions(); err != nil {
			ctx.StatusCode(500)
			ctx.Writef("error : %v\n", err)
			return
		}
		ctx.Writef("ok")
	})
	app.Run(iris.Addr(":8080"))
}
