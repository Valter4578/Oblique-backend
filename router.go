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

	// expenses
	router.HandleFunc("/expenses", expense.GetExpenses)
	router.HandleFunc("/expense/{id}", expense.GetExpense)
	router.HandleFunc("/expense", expense.AddExpense).Methods("POST")
	router.HandleFunc("/expenses/{id}", expense.UpdateExpenses).Methods("PUT")

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
