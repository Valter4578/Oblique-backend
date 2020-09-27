package db

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"oblique/iternal/app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertCategory(category *model.Category) *mongo.InsertOneResult {
	collection := DB.database().Collection(categories)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, category)
	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}

func InsertCategoryToWallet(category *model.Category, walletID primitive.ObjectID) error {
	collection := DB.database().Collection(categories)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, category)
	if err != nil {
		log.Println(err)
		return err
	}

	stringID := fmt.Sprintf("%v", result.InsertedID)
	i := strings.IndexByte(stringID, '"')
	if i != -1 {
		stringID = stringID[i+1 : len(stringID)-2]
	}

	id, err := primitive.ObjectIDFromHex(stringID)
	if err != nil {
		log.Println(err)
		return err
	}

	// addOperation(categoryID, id)
	addCategory(walletID, id)
	return nil
}

func GetCategory(id primitive.ObjectID) (*model.Category, error) {
	collection := DB.database().Collection(categories)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var category model.Category
	err := collection.FindOne(ctx, model.Category{ID: id}).Decode(&category)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &category, nil
}

func GetCategories() (*[]model.Category, error) {
	collection := DB.database().Collection(categories)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []model.Category
	for cursor.Next(ctx) {
		var category model.Category
		err = cursor.Decode(&category)
		if err != nil {
			log.Println(1)
			log.Println(err)
			return nil, err
		}

		categories = append(categories, category)
	}

	err = cursor.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &categories, nil
}

func UpdateCategory(id primitive.ObjectID, update bson.D) *mongo.UpdateResult {
	collection := DB.database().Collection(categories)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}

func addOperation(categoryID primitive.ObjectID, operationID primitive.ObjectID) {
	collection := DB.database().Collection(categories)
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
		log.Println(err)
		return
	}

	return
}
