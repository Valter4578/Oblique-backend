package db

import (
	"context"
	"log"
	"oblique/iternal/app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// GetUser gets email and password and searchs for user in db. Returns pointer to user and error if it's exist.
// Needs to sing in
func GetUser(email string, password string) (*model.User, error) {
	collection := DB.database().Collection(users)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	err := collection.FindOne(ctx, model.User{
		Email:    email,
		Password: password,
	}).Decode(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

// CreateUser gets pointer to user model. Returns error if it's exist
func CreateUser(user *model.User) error {
	collection := DB.database().Collection(users)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// GetUserDetails gets email and returns pointer to user's model and returns error if it's exist
func GetUserDetails(email string) (*model.User, error) {
	collection := DB.database().Collection(users)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}
