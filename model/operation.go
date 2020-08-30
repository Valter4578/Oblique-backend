package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Operation structure
type Operation struct {
	Title  string             `json:"title,omitempty" bson:"title,omitempty"`
	Amount int                `json:"amount,omitempty" bson:"amount,omitempty"`
	Type   string             `json:"type,omitempty" bson:"type,omitempty"`
	Time   time.Time          `json:"time,omitempty" bson:"time,omitempty"`
	ID     primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
}

// LastOperationID is identifier of last expense
var LastOperationID int
