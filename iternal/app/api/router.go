package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"oblique/iternal/app/api/endpoints"

	"github.com/gorilla/mux"
)

func Route() {
	log.Println("Starting routing")
	router := mux.NewRouter().StrictSlash(true)

	// operations
	router.HandleFunc("/operations", endpoints.GetOperations)
	router.HandleFunc("/operation/{id}", endpoints.DeleteOperation).Methods("DELETE")
	router.HandleFunc("/operation/{id}", endpoints.GetOperation)
	router.HandleFunc("/operation", endpoints.AddOperation).Methods("POST")

	// category
	router.HandleFunc("/categories", endpoints.GetAllCategories)
	router.HandleFunc("/category/{id}", endpoints.GetCategory)
	router.HandleFunc("/category", endpoints.AddCategory).Methods("POST")
	router.HandleFunc("/mostUsedCategories", endpoints.GetMostUsedCategories)
	// router.HandleFunc("/statistic", category.GetCategoriesStatistic)

	// wallet
	router.HandleFunc("/wallets", endpoints.GetWallets)
	router.HandleFunc("/wallet", endpoints.AddWallet).Methods("POST")
	router.HandleFunc("/wallet/{id}", endpoints.DeleteWallet).Methods("DELETE")
	router.HandleFunc("/wallet/{id}", endpoints.GetWallet)

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
