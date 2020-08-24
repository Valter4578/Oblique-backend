package database

import (
	"context"
	"oblique/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertWallet(wallet *model.Wallet) *mongo.InsertOneResult {
	collection := Client.Database("oblique").Collection("wallets")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, wallet)

	return result
}
