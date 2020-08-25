package main

import (
	"oblique/database"
	"os"
)

func main() {
	uri := os.Getenv("URI")
	database.ConnectDB(&uri)

	route()
}
