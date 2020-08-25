package database

import (
	"context"
	"log"
	"oblique/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertWallet gets wallet pointer and returns pointer to mongo.InsertOneResult
func InsertWallet(wallet *model.Wallet) *mongo.InsertOneResult {
	log.Println("Database: InsertWallet")
	collection := client.Database("oblique-dev").Collection("wallets")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, wallet)
	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}

// GetWallet gets id of wallet and pointer to wallet for decode result into it
func GetWallet(id primitive.ObjectID, wallet *model.Wallet) *[]byte {
	log.Println("Database: GetWallet")

	collection := client.Database("oblique-dev").Collection("wallets")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, model.Wallet{ID: id}).Decode(&wallet)
	if err != nil {
		msg := []byte(`{ "message": "` + err.Error() + `" }`)
		return &msg
	}

	return nil
}
