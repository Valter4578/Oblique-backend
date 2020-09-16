package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OperationType ...
type OperationType int

const (
	Expense OperationType = iota
	Revenue
)

func (o OperationType) String() string {
	return [...]string{"expense", "revenue"}[o]
}

// Operation structure
type Operation struct {
	Title  string             `json:"title,omitempty" bson:"title,omitempty"`
	Amount int                `json:"amount,omitempty" bson:"amount,omitempty"`
	Type   OperationType      `json:"type,omitempty" bson:"type,omitempty"`
	Time   time.Time          `json:"time,omitempty" bson:"time,omitempty"`
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}
