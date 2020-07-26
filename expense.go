package main

// Expense structure
type Expense struct {
	Title  string `json:"title,omitempty"`
	Amount int    `json:"amount,omitempty"`
	ID     int    `json:"id,omitempty"`
}

// Expenses ...
var Expenses []Expense

var LastExpenseID int
