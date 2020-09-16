package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	URI    string
	Client *mongo.Client
	DBName string
}

const (
	categories = "categories"
	wallets    = "wallets"
	operation  = "operations"
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

	log.Println("Connected to database")
}

func (d *Database) database() *mongo.Database {
	return d.Client.Database(d.DBName)
}

func (d *Database) insert(object interface{}, collectionName string) (*mongo.InsertOneResult, error) {
	collection := d.database().Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, object)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (d *Database) getOne(id primitive.ObjectID, collectionName string, object interface{}) interface{} {
	collection := d.database().Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var obj interface{}
	err := collection.FindOne(ctx, object).Decode(&obj)
	if err != nil {
		log.Println(err)
		return nil
	}

	return obj
}

// Client ...
var client *mongo.Client

// ConnectDB used to connect to MongoDB Atlas
func ConnectDB(uri string) {
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
