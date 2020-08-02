package model

// Category ...
type Category struct {
	Title      string `json:"title,omitempty" bson:"title,omitempty"`
	ImageName  string `json:"imageName,omitempty" bson:"imageName,omitempty"`
	Color      string `json:"color,omitempty" bson:"color,omitempty"`
	Operations []Operation
}

var Categories []Category = []Category{
	Category{Title: "Travel", ImageName: "travel", Color: "#fff", Operations: []Operation{
		Operation{Title: "Barcelona trip", Amount: 20000, Type: "expense", ID: 1},
		Operation{Title: "Paris trip", Amount: 351521, Type: "expense", ID: 2},
	}}, Category{Title: "Groceries", ImageName: "groceries", Color: "#652425", Operations: []Operation{
		Operation{Title: "Vegetables", Amount: 20, Type: "expense", ID: 3},
		Operation{Title: "Fruits", Amount: 10, Type: "expense", ID: 4},
		Operation{Title: "Meat", Amount: 40, Type: "expense", ID: 5},
	}},
}
