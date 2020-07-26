package main

// Expense structure
type Expense struct {
	Title  string `json:"title,omitempty" bson:"title,omitempty"`
	Amount int    `json:"amount,omitempty" bson:"amount,omitempty"`
	ID     int    `json:"id,omitempty" bson:"id,omitempty"`
}

// Expenses ...
var Expenses []Expense

// LastExpenseID is identifier of last expense
var LastExpenseID int
