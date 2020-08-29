package database

import (
	"context"
	"log"
	"oblique/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertCategory(category *model.Category) *mongo.InsertOneResult {
	log.Println("Database: InsertCategory")
	collection := client.Database("oblique-dev").Collection("categories")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, category)
	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}

func GetCategory(id primitive.ObjectID, category *model.Category) *error {
	log.Println("Database: GetCategory")

	collection := client.Database("oblique-dev").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, model.Category{ID: id}).Decode(&category)
	if err != nil {
		log.Println(err)
		return &err
	}

	return nil
}
