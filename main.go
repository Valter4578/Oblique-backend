package main

import (
	"oblique/db"
	"os"
)

func main() {
	uintptr := os.Getenv("URI")
	db.ConnectDB(uri)

	route()
}
