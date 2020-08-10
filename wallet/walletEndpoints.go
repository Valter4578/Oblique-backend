package wallet

import (
	"encoding/json"
	"log"
	"net/http"

	"oblique/model"
)

func getAllWallets(w http.ResponseWriter, r *http.Request) {
	log.Println("getAllWallets")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(model.Wallets)
}
