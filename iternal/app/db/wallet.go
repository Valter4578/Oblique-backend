package db

import (
	"context"
	"log"
	"time"

	"oblique/iternal/app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertWallet gets wallet pointer and returns pointer to mongo.InsertOneResult
func InsertWallet(wallet *model.Wallet) *mongo.InsertOneResult {
	log.Println("Database: InsertWallet")
	collection := DB.database().Collection(wallets)

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
func GetWallet(id primitive.ObjectID) (*model.Wallet, error) {
	log.Println("Database: GetWallet")

	collection := DB.database().Collection(wallets)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var wallet *model.Wallet
	err := collection.FindOne(ctx, model.Wallet{ID: id}).Decode(&wallet)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return wallet, nil
}

func GetWallets() (*[]model.Wallet, error) {
	log.Println("Database: GetWallets")

	collection := DB.database().Collection(wallets)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var wallets []model.Wallet
	for cursor.Next(ctx) {
		var wallet model.Wallet
		err = cursor.Decode(&wallet)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		wallets = append(wallets, wallet)
	}

	err = cursor.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &wallets, nil
}

func UpdateWallet(id primitive.ObjectID, update bson.D) *mongo.UpdateResult {
	log.Println("Database: UpdateWallet")

	collection := DB.database().Collection(wallets)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}

func DeleteWallet(id primitive.ObjectID) error {
	log.Println("Db: DeleteWallet")

	collection := DB.database().Collection(wallets)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}