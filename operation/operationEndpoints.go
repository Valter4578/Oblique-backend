package operation

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"oblique/database"
	"oblique/model"
)

// GetOperations is get method that returns all expenses
func GetOperations(w http.ResponseWriter, r *http.Request) {
	log.Println("GetExpenses")
	w.Header().Set("Content-Type", "application/json")

	// params := r.URL.Query()

	// TODO:- Add check of type
	// opType := params.Get("type")

	err, operations := database.GetOperations()
	if err != nil {
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(operations)
}

// GetOperation is get method that returns expense by id
func GetOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("GetExpense")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		log.Println(err)
		return
	}

	operation, err := database.GetOperation(id)
	if err != nil {
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(operation)
}

// AddOperation is post method for add new user's expense
func AddOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("AddOperation")
	w.Header().Set("Content-Type", "application/json")

	var operation model.Operation

	err := json.NewDecoder(r.Body).Decode(&operation)
	if err != nil {
		log.Println("Decode error: " + err.Error())
		return
	}

	result := database.InsertOperation(&operation)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)
}

// UpdateOperation is PUT method that updates expense by id
// func UpdateOperation(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	for _, item := range allOperations() {
// 		if item.ID == id {

// 			category, _ := category.FindCategory(item.Title)

// 			var operation model.Operation
// 			decoder := json.NewDecoder(r.Body)
// 			decoder.Decode(&operation)

// 			category.Operations[id] = operation

// 			json.NewEncoder(w).Encode(&operation)

// 			return
// 		}
// 	}

// 	json.NewEncoder(w).Encode(model.Categories)
// }

// // DeleteOperation is DELETE method that deletes expense by id
// func DeleteOperation(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	for indx, item := range allOperations() {
// 		if item.ID == id {
// 			category, _ := category.FindCategory(item.Title)
// 			category.Operations = append(category.Operations[:indx], category.Operations[indx+1:]...)

// 			break
// 		}
// 	}

// 	json.NewEncoder(w).Encode(model.Categories)
// }
