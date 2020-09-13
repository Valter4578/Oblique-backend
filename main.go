package main

import (
	"oblique/db"
	"os"
)

func main() {
	password := os.Getenv("DBPASSWORD")
	db.ConnectDB(password)

	route()
}
