package model

import "time"

// Expense structure
type Expense struct {
	Title    string `json:"title,omitempty" bson:"title,omitempty"`
	Amount   int    `json:"amount,omitempty" bson:"amount,omitempty"`
	Category Category
	Time     time.Time `json:"time,omitempty" bson:"time,omitempty"`
	ID       int       `json:"id,omitempty" bson:"id,omitempty"`
}

// mock data
// Expenses ...
var Expenses []Expense = []Expense{
	Expense{Title: "Barcelona trip", Amount: 20000, Category: Category{
		Title:     "Travel",
		ImageName: "Travel",
		Color:     "#fff",
	}, ID: 1},
	Expense{Title: "Groceries", Amount: 100, Category: Category{
		Title:     "Food",
		ImageName: "Food",
		Color:     "#5454454",
	}, ID: 2},
	Expense{Title: "Bike", Amount: 2202020, Category: Category{
		Title:     "Transport",
		ImageName: "Transport",
		Color:     "#56246GFW",
	}, ID: 3},
}

// LastExpenseID is identifier of last expense
var LastExpenseID int
