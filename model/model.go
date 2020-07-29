package model

// Expense structure
type Expense struct {
	Title  string `json:"title,omitempty" bson:"title,omitempty"`
	Amount int    `json:"amount,omitempty" bson:"amount,omitempty"`
	ID     int    `json:"id,omitempty" bson:"id,omitempty"`
}

// Expenses ...
var Expenses []Expense = []Expense{
	Expense{Title: "Barcelona trip", Amount: 20000, ID: 1},
	Expense{Title: "Groceries", Amount: 100, ID: 2},
	Expense{Title: "Bike", Amount: 2202020, ID: 3},
}

// LastExpenseID is identifier of last expense
var LastExpenseID int
