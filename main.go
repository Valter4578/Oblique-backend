package main

import (
	"oblique/db"
	"os"
)

func main() {
	uri := os.Getenv("URI")
	db.ConnectDB(uri)

	route()
}
