package expense

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"mellow/category"
	"mellow/model"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

// GetExpenses is get method that returns all expenses
func GetExpenses(w http.ResponseWriter, r *http.Request) {
	log.Println("GetExpenses")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(allExpenses())
}

// GetExpense is get method that returns expense by id
func GetExpense(w http.ResponseWriter, r *http.Request) {
	log.Println("GetExpense")

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		return
	}

	for _, expense := range allExpenses() {
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
	expense.CategoryTitle = params.Get("categoryTitle")

	allExpenses := allExpenses()
	expense.ID = allExpenses[len(allExpenses)-1].ID + 1

	expense.Time = time.Now()

	// // model.Expenses = append(model.Expenses, expense)
	// for _, category := range model.Categories {
	// 	if strings.ToLower(category.Title) == strings.ToLower(expense.Title) {
	// 		category.Expenses = append(category.Expenses, expense)
	// 	}
	// }

	category, index := category.FindCategory(expense.Title)
	category.Expenses = append(category.Expenses, expense)

	model.Categories[index] = category

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(model.Categories)
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

	for _, item := range allExpenses() {
		if item.ID == id {

			category, _ := category.FindCategory(item.Title)

			// model.Expenses = append(model.Expenses[:indx], model.Expenses[indx+1:]...)

			var expense model.Expense
			decoder := json.NewDecoder(r.Body)
			decoder.Decode(&expense)

			// model.Expenses = append(model.Expenses, expense)
			category.Expenses[id] = expense

			json.NewEncoder(w).Encode(&expense)

			return
		}
	}

	json.NewEncoder(w).Encode(model.Categories)
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

	for indx, item := range allExpenses() {
		if item.ID == id {
			category, _ := category.FindCategory(item.Title)
			category.Expenses = append(category.Expenses[:indx], category.Expenses[indx+1:]...)

			break
		}
	}

	json.NewEncoder(w).Encode(model.Categories)
}

func GetExpenseByCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("GetExpenseByCategory")

	w.Header().Set("Content-Type", "application/json")

	// vars := mux.Vars(r)
	// title := vars["title"]

}
