package model

import "time"

// Operation structure
type Operation struct {
	Title  string    `json:"title,omitempty" bson:"title,omitempty"`
	Amount int       `json:"amount,omitempty" bson:"amount,omitempty"`
	Type   string    `json:"type,omitempty" bson:"type,omitempty"`
	Time   time.Time `json:"time,omitempty" bson:"time,omitempty"`
	ID     int       `json:"id,omitempty" bson:"id,omitempty"`
}

// LastOperationID is identifier of last expense
var LastOperationID int
