package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"mellow/expense"
)

func route() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/expenses", expense.GetExpenses)
	router.HandleFunc("/expenses/{id}", expense.GetExpense)
	router.HandleFunc("/expenses", expense.AddExpense).Methods("POST")
	router.HandleFunc("/expenses/{id}", expense.UpdateExpenses).Methods("PUT")

	http.ListenAndServe(":8080", router)
}
