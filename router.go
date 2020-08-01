package main

import (
	"log"
	"net/http"

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

	log.Fatal(http.ListenAndServe("$PORT", router))
}
