package wallet

import (
	"encoding/json"
	"log"
	"net/http"

	"oblique/model"

	"github.com/gorilla/mux"
)

func GetAllWallets(w http.ResponseWriter, r *http.Request) {
	log.Println("getAllWallets")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(model.Wallets)
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	log.Println("getWallet")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	title := params["title"]

	wallet, _ := FindWallet(title)

	json.NewEncoder(w).Encode(wallet)
}
