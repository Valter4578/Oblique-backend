package model

// Category ...
type Category struct {
	Title     string `json:"title,omitempty" bson:"title,omitempty"`
	ImageName string `json:"imageName,omitempty" bson:"imageName,omitempty"`
	Color     string `json:"color,omitempty" bson:"color,omitempty"`
	Expenses  []Expense
}

var Categories []Category = []Category{
	Category{Title: "Travel", ImageName: "travel", Color: "#fff", Expenses: []Expense{
		Expense{Title: "Barcelona trip", Amount: 20000, ID: 1},
		Expense{Title: "Paris trip", Amount: 351521, ID: 2},
	}}, Category{Title: "Groceries", ImageName: "groceries", Color: "#652425", Expenses: []Expense{
		Expense{Title: "Vegetables", Amount: 20, ID: 3},
		Expense{Title: "Fruits", Amount: 10, ID: 4},
		Expense{Title: "Meat", Amount: 40, ID: 5},
	}},
}
