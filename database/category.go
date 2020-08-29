package database

import (
	"context"
	"log"
	"oblique/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func GetCategories(categories *[]model.Category) *error {
	log.Println("Database: GetCategories")

	collection := client.Database("oblique-dev").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return &err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var category model.Category
		err = cursor.Decode(&categories)
		if err != nil {
			log.Println(err)
			return &err
		}

		*categories = append(*categories, category)
	}

	err = cursor.Err()
	if err != nil {
		log.Println(err)
		return &err
	}

	return nil
}
