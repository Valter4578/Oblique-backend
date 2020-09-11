package db

import (
	"context"
	"log"
	"oblique/logger"
	"oblique/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertCategory is function that gets the pointer to the category structure and returns *mongo.InsertOneResult
func InsertCategory(category *model.Category) *mongo.InsertOneResult {
	log.Println("Database: InsertCategory")
	collection := client.Database("oblique-dev").Collection("categories")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, category)
	if err != nil {
		logger.LogError(&err)
		return nil
	}

	return result
}

<<<<<<< HEAD:database/category.go
// GetCategory is function that gets the id of operation and returns error if it exist and returns *model.Category
func GetCategory(id primitive.ObjectID) (error, *model.Category) {
=======
func GetCategory(id primitive.ObjectID, category *model.Category) error {
>>>>>>> master:db/category.go
	log.Println("Database: GetCategory")

	collection := client.Database("oblique-dev").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var category *model.Category
	err := collection.FindOne(ctx, model.Category{ID: id}).Decode(&category)
	if err != nil {
<<<<<<< HEAD:database/category.go
		log.Println(err)
		return err, nil
=======
		logger.LogError(&err)
		return err
>>>>>>> master:db/category.go
	}

	return nil, category
}

<<<<<<< HEAD:database/category.go
// GetCategories is function that returns error and returns all categories slice from db
func GetCategories() (error, *[]model.Category) {
=======
func GetCategories(categories *[]model.Category) error {
>>>>>>> master:db/category.go
	log.Println("Database: GetCategories")

	collection := client.Database("oblique-dev").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
<<<<<<< HEAD:database/category.go
		log.Println(err)
		return err, nil
=======
		logger.LogError(&err)
		return err
>>>>>>> master:db/category.go
	}
	defer cursor.Close(ctx)

	var categories *[]model.Category
	for cursor.Next(ctx) {
		var category model.Category
		err = cursor.Decode(&categories)
		if err != nil {
<<<<<<< HEAD:database/category.go
			log.Println(err)
			return err, nil
=======
			logger.LogError(&err)
			return err
>>>>>>> master:db/category.go
		}

		*categories = append(*categories, category)
	}

	err = cursor.Err()
	if err != nil {
<<<<<<< HEAD:database/category.go
		log.Println(err)
		return err, nil
=======
		logger.LogError(&err)
		return err
>>>>>>> master:db/category.go
	}

	return nil, categories
}

// UpdateCategory  is function that gets the id and bson object. Returns *mongo.UpdateResult
func UpdateCategory(id primitive.ObjectID, update bson.D) *mongo.UpdateResult {
	log.Println("Database: UpdateCategory")

	collection := client.Database("oblique-dev").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		logger.LogError(&err)
		return nil
	}

	return result
}

func addOperation(categoryID primitive.ObjectID, operationID primitive.ObjectID) error {
	collection := client.Database("oblique-dev").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": categoryID}
	update := bson.D{
		{"$push", bson.D{
			{"operations", operationID},
		}},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.LogError(&err)
		return err
	}

	return nil
}
