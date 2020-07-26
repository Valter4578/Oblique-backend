package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func route() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/getExpenses", getExpenses)
	router.HandleFunc("/getExpenses/{id}", getExpense)
	router.HandleFunc("/addExpense", addExpense).Methods("POST")

	http.ListenAndServe(":8080", router)
}
