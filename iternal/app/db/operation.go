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

func InsertOperation(operation *model.Operation) *mongo.InsertOneResult {
	log.Println("Database: InsertOperation")
	collection := DB.database().Collection(operations)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, operation)
	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}

func InsertOperationToCategory(operation *model.Operation, categoryID primitive.ObjectID) error {
	log.Println("Database: InsertOperationToCategory")
	collection := DB.database().Collection(operations)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, operation)
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

	addOperation(categoryID, id)
	return nil
}

func GetOperation(id primitive.ObjectID) (*model.Operation, error) {
	log.Println("Database: GetOperation")

	collection := DB.database().Collection(operations)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var operation model.Operation
	err := collection.FindOne(ctx, model.Operation{ID: id}).Decode(&operation)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &operation, nil
}

func GetOperations() (*[]model.Operation, error) {
	log.Println("Database: GetOperations")

	collection := DB.database().Collection(operations)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var operations []model.Operation
	for cursor.Next(ctx) {
		var operation model.Operation
		err = cursor.Decode(&operation)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		operations = append(operations, operation)
	}

	err = cursor.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &operations, nil
}

func UpdateOperation(id primitive.ObjectID, update bson.D) *mongo.UpdateResult {
	log.Println("Database: UpdateOperation")

	collection := DB.database().Collection(operations)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}

func DeleteOperation(id primitive.ObjectID) error {
	log.Println("Database: DeleteOperation")

	collection := DB.database().Collection(operations)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
