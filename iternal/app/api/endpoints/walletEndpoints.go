package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"oblique/iternal/app/db"
	"oblique/iternal/app/logger"
	"oblique/iternal/app/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetWallets(w http.ResponseWriter, r *http.Request) {
	log.Println("GetWallets")
	w.Header().Set("Content-Type", "application/json")

	wallets, err := db.GetWallets()
	if err != nil {
		w.Write([]byte(logger.JSONError(err)))
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(&wallets)
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	log.Println("getWallet")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println(err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	wallet, err := db.GetWallet(id)
	if err != nil {
		log.Println(err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	json.NewEncoder(w).Encode(&wallet)
}

func AddWallet(w http.ResponseWriter, r *http.Request) {
	log.Println("AddWallet")
	w.Header().Set("Content-Type", "application/json")

	var wallet model.Wallet
	err := json.NewDecoder(r.Body).Decode(&wallet)
	if err != nil {
		log.Println(err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	result := db.InsertWallet(&wallet)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteWallet")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println(err)
		return
	}

	err = db.DeleteWallet(id)
	if err != nil {
		json := logger.JSONError(err)
		w.Write([]byte(json))
		return
	}

	msg := logger.JSONMessage("Successfully delete operation with id ", id.String)
	w.Write([]byte(msg))
}
