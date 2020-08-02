package model

import "time"

// Expense structure
type Expense struct {
	Title  string    `json:"title,omitempty" bson:"title,omitempty"`
	Amount int       `json:"amount,omitempty" bson:"amount,omitempty"`
	Time   time.Time `json:"time,omitempty" bson:"time,omitempty"`
	ID     int       `json:"id,omitempty" bson:"id,omitempty"`
}

// LastExpenseID is identifier of last expense
var LastExpenseID int
