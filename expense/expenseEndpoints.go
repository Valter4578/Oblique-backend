package expense

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"./model"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func getExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Expenses)
}

func getExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	// id := vars["id"]
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		return
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

	expense.ID = Expenses[len(Expenses)-1].ID + 1

	Expenses = append(Expenses, expense)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(Expenses)
}

func updateExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		return
	}

	for indx, item := range Expenses {
		if item.ID == id {
			Expenses = append(Expenses[:indx], Expenses[indx+1:]...)

			var expense Expense
			decoder := json.NewDecoder(r.Body)
			decoder.Decode(&expense)

			Expenses = append(Expenses, expense)

			json.NewEncoder(w).Encode(&expense)

			return
		}
	}

	json.NewEncoder(w).Encode(Expenses)
}

func deleteExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		return
	}

	for indx, item := range Expenses {
		if item.ID == id {
			Expenses = append(Expenses[:indx], Expenses[indx+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(Expenses)
}
