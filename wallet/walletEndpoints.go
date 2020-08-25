package wallet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"oblique/database"
	"oblique/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllWallets(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllWallets")
	w.Header().Set("Content-Type", "application/json")

	var wallets []model.Wallet
	err := database.GetWallets(&wallets)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(wallets)
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	log.Println("getWallet")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, `{"message":"%v"}`, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var wallet model.Wallet

	msg := database.GetWallet(id, &wallet)
	if msg != nil {
		log.Println(msg)
		w.Write(*msg)
		return
	}

	json.NewEncoder(w).Encode(wallet)
}

func AddWallet(w http.ResponseWriter, r *http.Request) {
	log.Println("AddWallet")
	w.Header().Set("Content-Type", "application/json")

	// _ := r.URL.Query()
	var wallet model.Wallet
	err := json.NewDecoder(r.Body).Decode(&wallet)
	if err != nil {
		log.Println("Decode error: " + err.Error())
		return
	}
	result := database.InsertWallet(&wallet)

	log.Println(result)
	json.NewEncoder(w).Encode(result)

	w.WriteHeader(http.StatusCreated)
}
