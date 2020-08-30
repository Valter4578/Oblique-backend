package database

import (
	"context"
	"log"
	"oblique/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOperation(operation *model.Operation) *mongo.InsertOneResult {
	log.Println("Database: InsertOperation")
	collection := client.Database("oblique-dev").Collection("operations")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, operation)
	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}

func GetOperation(id primitive.ObjectID, operation *model.Operation) *error {
	log.Println("Database: GetOperation")

	collection := client.Database("oblique-dev").Collection("operations")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, model.Operation{ID: id}).Decode(&operation)
	if err != nil {
		log.Println(err)
		return &err
	}

	return nil
}
