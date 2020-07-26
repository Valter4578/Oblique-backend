package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	// mock data
	Expenses = []Expense{
		Expense{Title: "Barcelona trip", Amount: 20000, ID: 1},
		Expense{Title: "Groceries", Amount: 100, ID: 2},
		Expense{Title: "Bike", Amount: 2202020, ID: 3},
	}

	route()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pisya")
}

func getExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Expenses)
}

func getExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	// id := vars["id"]
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	for _, expense := range Expenses {
		if expense.ID == id {
			json.NewEncoder(w).Encode(expense)
		}
	}
}

func addExpense(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)

	LastExpenseID++

	var expense Expense
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&expense)
	if err != nil {
		log.Fatal(err)
	}

	expense.ID = LastExpenseID

	Expenses = append(Expenses, expense)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(Expenses)
}
