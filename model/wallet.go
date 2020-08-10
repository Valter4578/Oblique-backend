package model

type Wallet struct {
	Title      string   `json:"title,omitempty" bson:"title,omitempty"`
	Colors     []string `json:"colors,omitempty" bson:"colors,omitempty"`
	Categories []Category
}

var Wallets = Wallet{
	Title:  "Main wallet",
	Colors: []string{"#2924FF", "#FF00E5"},
	Categories: []Category{
		Category{Title: "Travel", ImageName: "travel", Color: "#fff", Operations: []Operation{
			Operation{Title: "Barcelona trip", Amount: 20000, Type: "expense", ID: 1},
			Operation{Title: "Paris trip", Amount: 351521, Type: "expense", ID: 2},
		}}, Category{Title: "Groceries", ImageName: "groceries", Color: "#652425", Operations: []Operation{
			Operation{Title: "Vegetables", Amount: 20, Type: "expense", ID: 3},
			Operation{Title: "Fruits", Amount: 10, Type: "expense", ID: 4},
			Operation{Title: "Meat", Amount: 40, Type: "expense", ID: 5},
		}},
	},
}
