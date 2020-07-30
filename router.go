package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"mellow/expense"
)

func route() {
	log.Println("Starting routing")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/expenses", expense.GetExpenses)
	router.HandleFunc("/expenses/{id}", expense.GetExpense)
	router.HandleFunc("/addExpense", expense.AddExpense).Methods("POST")
	router.HandleFunc("/expenses/{id}", expense.UpdateExpenses).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
	log.Println("Start listening at port :8080")

}
