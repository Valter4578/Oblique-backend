package operation

import (
	"encoding/json"
	"log"
	"net/http"
<<<<<<< HEAD
=======
	"oblique/db"
	"oblique/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
>>>>>>> master

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

<<<<<<< HEAD
	"oblique/database"
=======
>>>>>>> master
	"oblique/model"
)

// GetOperations is get method that returns all expenses
func GetOperations(w http.ResponseWriter, r *http.Request) {
	log.Println("GetExpenses")
	w.Header().Set("Content-Type", "application/json")

<<<<<<< HEAD
	// params := r.URL.Query()

	// TODO:- Add check of type
	// opType := params.Get("type")

	err, operations := database.GetOperations()
	if err != nil {
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(operations)
=======
	var operations []model.Operation
	err := db.GetOperations(&operations)
	if err != nil {
		logger.LogError(&err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	json.NewEncoder(w).Encode(&operations)
>>>>>>> master
}

// GetOperation is get method that returns expense by id
func GetOperation(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	log.Println("GetExpense")
=======
	log.Println("GetOperation")

>>>>>>> master
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		logger.LogError(&err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

<<<<<<< HEAD
	operation, err := database.GetOperation(id)
	if err != nil {
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(operation)
=======
	var operation model.Operation
	err = db.GetOperation(id, &operation)
	if err != nil {
		logger.LogError(&err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	json.NewEncoder(w).Encode(&operation)
>>>>>>> master
}

// AddOperation is post method for add new user's expense
func AddOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("AddOperation")
<<<<<<< HEAD
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
=======

	var operation model.Operation
	err := json.NewDecoder(r.Body).Decode(&operation)
	if err != nil {
		logger.LogError(&err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	params := r.URL.Query()
	id := params.Get("categoryId")
	if id != "" {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logger.LogError(&err)
			w.Write([]byte(logger.JSONError(err)))
			return
		}
		err = db.InsertOperationToCategory(&operation, objID)
		if err != nil { 
			logger.LogError(&err)
			w.Write([]byte(logger))
		}
	} else {
		result := db.InsertOperation(&operation)
		json.NewEncoder(w).Encode(result)
	}

}

// DeleteOperation is DELETE method that deletes expense by id
>>>>>>> master
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
