package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"oblique/category"
	"oblique/operation"
	"oblique/wallet"
)

func route() {
	log.Println("Starting routing")
	router := mux.NewRouter().StrictSlash(true)

	// operations
	router.HandleFunc("/operations", operation.GetOperations)
	router.HandleFunc("/operation/{id}", operation.GetOperation)
	router.HandleFunc("/operation", operation.AddOperation).Methods("POST")
	// router.HandleFunc("/operation/{id}", operation.UpdateOperation).Methods("PUT")

	// category
	router.HandleFunc("/categories", category.GetAllCategories)
	router.HandleFunc("/category/{id}", category.GetCategory)
	router.HandleFunc("/category", category.AddCategory).Methods("POST")
	router.HandleFunc("/mostUsedCategories", category.GetMostUsedCategories)
	router.HandleFunc("/statistic", category.GetCategoriesStatistic)

	// wallet
	router.HandleFunc("/wallets", wallet.GetAllWallets)
	router.HandleFunc("/wallet", wallet.AddWallet).Methods("POST")
	router.HandleFunc("/wallet/{id}", wallet.GetWallet)

	log.Fatal(http.ListenAndServe(getPort(), router))
}

func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	return ":" + port
}
