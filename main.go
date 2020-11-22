package main

import (
	"flag"
	"oblique/iternal/app/api"
	"oblique/iternal/app/db"

	"os"
)

var isDebugFlag = flag.Bool("debug", false, "Defines if application will start in debug mode")

func main() {
	flag.Parse()

	uri := os.Getenv("URI")
	db := &db.Database{
		URI:    uri,
		DBName: "oblique-dev",
	}

	db.Connect()

	api := &api.API{
		IsDebugMode: *isDebugFlag,
	}

	api.Route()
}
