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

// GetOperations is get method that returns all expenses
func GetOperations(w http.ResponseWriter, r *http.Request) {
	log.Println("GetExpenses")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(allOperations())
}

// GetOperation is get method that returns expense by id
func GetOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("GetExpense")

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		return
	}

	for _, operation := range allOperations() {
		if operation.ID == id {
			json.NewEncoder(w).Encode(operation)
		}
	}
}

// AddOperation is post method for add new user's expense
func AddOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("addExpense")
	log.Println(r.URL.Query())

	params := r.URL.Query()

	model.LastOperationID++

	var operation model.Operation

	operation.Title = params.Get("title")
	operation.Amount, _ = strconv.Atoi(params.Get("amount"))
	operation.Type = params.Get("type")

	allOperations := allOperations()
	operation.ID = allOperations[len(allOperations)-1].ID + 1

	operation.Time = time.Now()

	category, index := category.FindCategory(operation.Title)
	category.Operations = append(category.Operations, operation)

	model.Categories[index] = category

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(model.Categories)
}

// UpdateOperation is PUT method that updates expense by id
func UpdateOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range allOperations() {
		if item.ID == id {

			category, _ := category.FindCategory(item.Title)

			var operation model.Operation
			decoder := json.NewDecoder(r.Body)
			decoder.Decode(&operation)

			category.Operations[id] = operation

			json.NewEncoder(w).Encode(&operation)

			return
		}
	}

	json.NewEncoder(w).Encode(model.Categories)
}

// DeleteOperation is DELETE method that deletes expense by id
func DeleteOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		return
	}

	for indx, item := range allOperations() {
		if item.ID == id {
			category, _ := category.FindCategory(item.Title)
			category.Operations = append(category.Operations[:indx], category.Operations[indx+1:]...)

			break
		}
	}

	json.NewEncoder(w).Encode(model.Categories)
}
