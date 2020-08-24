package main

import (
	"oblique/database"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	uri := os.Getenv("URI")
	database.ConnectDB(&uri)

	route()
}
