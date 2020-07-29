package main

import (
	"net/http"

	"github.com/gorilla/mux"
	
	"expense"
)

func route() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", expense.homeHandler)
	router.HandleFunc("/expenses", expense.getExpenses)
	router.HandleFunc("/expenses/{id}", expense.getExpense)
	router.HandleFunc("/expenses", addExpense).Methods("POST")
	router.HandleFunc("/expenses/{id}", updateExpenses).Methods("PUT")

	http.ListenAndServe(":8080", router)
}
