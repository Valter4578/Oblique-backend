package operation

import (
	"encoding/json"
	"log"
	"net/http"
	"oblique/db"
	"oblique/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"

	"oblique/model"
)

// GetOperations is get method that returns all expenses
func GetOperations(w http.ResponseWriter, r *http.Request) {
	log.Println("GetExpenses")

	w.Header().Set("Content-Type", "application/json")

	var operations []model.Operation
	err := db.GetOperations(&operations)
	if err != nil {
		logger.LogError(&err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	json.NewEncoder(w).Encode(&operations)
}

// GetOperation is get method that returns expense by id
func GetOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("GetOperation")

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		logger.LogError(&err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	var operation model.Operation
	err = db.GetOperation(id, &operation)
	if err != nil {
		logger.LogError(&err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	json.NewEncoder(w).Encode(&operation)
}

// AddOperation is post method for add new user's expense
func AddOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("AddOperation")

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
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = db.InsertOperationToCategory(&operation, objID)
		if err != nil {
			logger.LogError(&err)
			w.Write([]byte(logger.JSONError(err)))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		msg := logger.JSONMessage("The operation was successfully added to the category")
		w.Write([]byte(msg))

		w.WriteHeader(http.StatusCreated)
	} else {
		result := db.InsertOperation(&operation)
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(result)
	}
}

// DeleteOperation is DELETE method that deletes expense by id
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
