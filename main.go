package main

import (
	"oblique/iternal/app/api"
	"oblique/iternal/app/db"

	"os"
)

func main() {
	uri := os.Getenv("URI")
	db := &db.Database{
		URI:    uri,
		DBName: "oblique-dev",
	}

	db.Connect()

	api.Route()
}
