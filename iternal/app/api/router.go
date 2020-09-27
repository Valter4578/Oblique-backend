package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"oblique/iternal/app/api/auth"
	"oblique/iternal/app/api/endpoints"

	"github.com/gorilla/mux"
)

func Route() {
	log.Println("Starting routing")
	router := mux.NewRouter().StrictSlash(true)

	// auth
	router.HandleFunc("/signin", Chain(auth.SignIn, Method("POST"), Logging())).Methods("POST")
	router.HandleFunc("/signup", Chain(auth.SignUp, Method("POST"), Logging())).Methods("POST")
	router.HandleFunc("/userDetails", Chain(auth.GetUserDetails, Method("GET"), Logging()))

	// operations
	router.HandleFunc("/operations", Chain(endpoints.GetOperations, Method("GET"), Logging()))
	router.HandleFunc("/operation/{id}", Chain(endpoints.DeleteOperation, Method("DELETE"), Logging())).Methods("DELETE")
	router.HandleFunc("/operation/{id}", Chain(endpoints.GetOperation, Method("GET"), Logging()))
	router.HandleFunc("/operation", Chain(endpoints.AddOperation, Method("POST"), Logging())).Methods("POST")

	// category
	router.HandleFunc("/categories", Chain(endpoints.GetAllCategories, Method("GET"), Logging()))
	router.HandleFunc("/category/{id}", Chain(endpoints.GetCategory, Method("GET"), Logging()))
	router.HandleFunc("/category", Chain(endpoints.AddCategory, Method("POST"), Logging())).Methods("POST")
	router.HandleFunc("/mostUsedCategories", Chain(endpoints.GetMostUsedCategories, Method("GET"), Logging()))
	// router.HandleFunc("/statistic", category.GetCategoriesStatistic)

	// wallet
	router.HandleFunc("/wallets", Chain(endpoints.GetWallets, Method("GET"), Logging()))
	router.HandleFunc("/wallet", Chain(endpoints.AddWallet, Method("POST"), Logging())).Methods("POST")
	router.HandleFunc("/wallet/{id}", Chain(endpoints.DeleteWallet, Method("DELETE"), Logging())).Methods("DELETE")
	router.HandleFunc("/wallet/{id}", Chain(endpoints.GetWallet, Method("GET"), Logging()))

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
