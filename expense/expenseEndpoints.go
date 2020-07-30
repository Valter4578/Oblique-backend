package expense

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"mellow/model"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

// GetExpenses is get method that returns all expenses
func GetExpenses(w http.ResponseWriter, r *http.Request) {
	log.Println("GetExpenses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Expenses)
}

// GetExpense is get method that returns expense by id
func GetExpense(w http.ResponseWriter, r *http.Request) {
	log.Println("GetExpense")

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	// id := vars["id"]
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		return
	}

	for _, expense := range model.Expenses {
		if expense.ID == id {
			json.NewEncoder(w).Encode(expense)
		}
	}
}

// AddExpense is post method for add new user's expense
func AddExpense(w http.ResponseWriter, r *http.Request) {
	log.Println("addExpense")
	log.Println(r.URL.Query())

	params := r.URL.Query()

	model.LastExpenseID++

	var expense model.Expense

	expense.Title = params.Get("title")
	expense.Amount, _ = strconv.Atoi(params.Get("amount"))
	expense.ID = model.Expenses[len(model.Expenses)-1].ID + 1
	expense.Time = time.Now()

	model.Expenses = append(model.Expenses, expense)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(model.Expenses)
}

// UpdateExpenses is PUT method that updates expense by id
func UpdateExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		return
	}

	for indx, item := range model.Expenses {
		if item.ID == id {
			model.Expenses = append(model.Expenses[:indx], model.Expenses[indx+1:]...)

			var expense model.Expense
			decoder := json.NewDecoder(r.Body)
			decoder.Decode(&expense)

			model.Expenses = append(model.Expenses, expense)

			json.NewEncoder(w).Encode(&expense)

			return
		}
	}

	json.NewEncoder(w).Encode(model.Expenses)
}

// DeleteExpense is DELETE method that deletes expense by id
func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		return
	}

	for indx, item := range model.Expenses {
		if item.ID == id {
			model.Expenses = append(model.Expenses[:indx], model.Expenses[indx+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(model.Expenses)
}
