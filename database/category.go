package database

import (
	"context"
	"log"
	"oblique/model"
	"time"

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
