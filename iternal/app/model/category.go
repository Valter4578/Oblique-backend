package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Category ...
type Category struct {
	Title      string               `json:"title,omitempty" bson:"title,omitempty"`
	ImageName  string               `json:"imageName,omitempty" bson:"imageName,omitempty"`
	Color      string               `json:"color,omitempty" bson:"color,omitempty"`
	Percantage float32              `json:"percantage,omitempty" bson:"percantage,omitempty"`
	Operations []primitive.ObjectID `json:"operations,omitempty" bson:"operations,omitempty"`
	ID         primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
}
