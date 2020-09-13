package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client ...
var client *mongo.Client

// ConnectDB used to connect to MongoDB Atlas
func ConnectDB(password string) {
	part := "mongodb+srv://valter:"
	uri := fmt.Sprintf("%v%v@oblique.bifuo.mongodb.net/oblique?retryWrites=true&w=majority", part, password)
	fmt.Println(uri)

	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")
}
