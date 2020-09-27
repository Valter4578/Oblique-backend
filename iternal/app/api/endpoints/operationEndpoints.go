package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"oblique/iternal/app/db"
	"oblique/iternal/app/logger"
	"oblique/iternal/app/model"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
)

// GetOperations is get method that returns all expenses
func GetOperations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	operations, err := db.GetOperations()
	if err != nil {
		w.Write([]byte(logger.JSONError(err)))
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(&operations)
}

// GetOperation is get method that returns expense by id
func GetOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		log.Println(err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	operation, err := db.GetOperation(id)
	if err != nil {
		w.Write([]byte(logger.JSONError(err)))
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(&operation)
}

// AddOperation is post method for add new user's expense
func AddOperation(w http.ResponseWriter, r *http.Request) {
	var operation model.Operation
	err := json.NewDecoder(r.Body).Decode(&operation)
	if err != nil {
		log.Println(err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	operation.Time = time.Now()

	params := r.URL.Query()
	id := params.Get("categoryId")
	if id != "" {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Println(err)
			w.Write([]byte(logger.JSONError(err)))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = db.InsertOperationToCategory(&operation, objID)
		if err != nil {
			log.Println(err)
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
func DeleteOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println(err)
		return
	}

	err = db.DeleteOperation(id)
	if err != nil {
		json := logger.JSONError(err)
		w.Write([]byte(json))
		return
	}

	msg := logger.JSONMessage("Successfully delete operation with id ", id.String)
	w.Write([]byte(msg))
}
