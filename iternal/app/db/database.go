package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	URI    string
	Client *mongo.Client
	DBName string
}

var DB *Database

const (
	categories = "categories"
	wallets    = "wallets"
	operations = "operations"
)

func (d *Database) Connect() {
	var err error
	d.Client, err = mongo.NewClient(options.Client().ApplyURI(d.URI))
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = d.Client.Connect(ctx)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	err = d.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	DB = d
	log.Println("Connected to database")
}

func (d *Database) database() *mongo.Database {
	return d.Client.Database(d.DBName)
}
