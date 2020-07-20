package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// mock data
	Expenses = []Expense{
		Expense{Title: "Barcelona trip", Amount: 20000, Id: "1"},
		Expense{Title: "Groceries", Amount: 100, Id: "2"},
		Expense{Title: "Bike", Amount: 2202020, Id: "3"},
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/getExpenses", getExpenses)
	router.HandleFunc("/getExpenses/{id}", getExpense)

	http.ListenAndServe(":8080", router)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pisya")
}

func getExpenses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Expenses)
}

func getExpense(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, expense := range Expenses {
		if expense.Id == id {
			json.NewEncoder(w).Encode(expense)
		}
	}
}
