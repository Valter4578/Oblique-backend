package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Wallet ...
type Wallet struct {
	Title      string             `json:"title,omitempty" bson:"title,omitempty"`
	Colors     []string           `json:"colors,omitempty" bson:"colors,omitempty"`
	Categories []Category         `json:"categories,omitempty" bson:"categories,omitempty"`
	Amount     int64              `json:"amount,omitempty" bson:amount,omitempty`
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}
