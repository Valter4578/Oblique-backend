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

// InsertOperation is function that gets pointer to Operation structure and returns pointer to mongo.InsertOneResult
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

// GetOperation is function that gets the id of operation and returns error if it exist and Operation structure
func GetOperation(id primitive.ObjectID) (*model.Operation, error) {
	log.Println("Database: GetOperation")

	collection := client.Database("oblique-dev").Collection("operations")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var operation *model.Operation
	err := collection.FindOne(ctx, model.Operation{ID: id}).Decode(&operation)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return operation, err
}

// GetOperations is method that returns error if it exists and returns slice of operations struct
func GetOperations() (*error, *[]model.Operation) {
	log.Println("Database: GetOPerations")

	collection := client.Database("oblique-dev").Collection("operations")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return &err, nil
	}
	defer cursor.Close(ctx)

	var operations *[]model.Operation
	for cursor.Next(ctx) {
		var operation model.Operation

		err = cursor.Decode(&operations)
		if err != nil {
			log.Println(err)
			return &err, nil
		}

		*operations = append(*operations, operation)
	}

	err = cursor.Err()
	if err != nil {
		log.Println(err)
		return &err, nil
	}

	return nil, operations
}

// UpdateOperation is function that gets the id of operation and gets bson object that contains information about the operation's update. Function returns pointer to mongo.UpdateResult
func UpdateOperation(id primitive.ObjectID, update bson.D) *mongo.UpdateResult {
	log.Println("Database: UpdateOperation")

	collection := client.Database("oblique-dev").Collection("operations")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}
