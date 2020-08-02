package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"mellow/category"
	"mellow/expense"
)

func route() {
	log.Println("Starting routing")
	router := mux.NewRouter().StrictSlash(true)

	// operations
	router.HandleFunc("/operations", expense.GetExpenses)
	router.HandleFunc("/operation/{id}", expense.GetExpense)
	router.HandleFunc("/operation", expense.AddExpense).Methods("POST")
	router.HandleFunc("/operation/{id}", expense.UpdateExpenses).Methods("PUT")

	// category
	router.HandleFunc("/categories", category.GetAllCategories)
	router.HandleFunc("/category/{title}", category.GetCategory)
	router.HandleFunc("/category", category.AddCategory).Methods("POST")
	router.HandleFunc("/mostUsedCategories", category.GetMostUsedCategories)

	log.Fatal(http.ListenAndServe(getPort(), router))
}

func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
