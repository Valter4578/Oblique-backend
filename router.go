package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func route() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/expenses", getExpenses)
	router.HandleFunc("/expenses/{id}", getExpense)
	router.HandleFunc("/expenses", addExpense).Methods("POST")
	router.HandleFunc("/expenses/{id}", updateExpenses).Methods("PUT")

	http.ListenAndServe(":8080", router)
}
