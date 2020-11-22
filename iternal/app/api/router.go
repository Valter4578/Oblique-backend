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

// Route create endpoints routes
func (api *API) Route() {
	log.Println("Starting routing")
	api.Router = mux.NewRouter().StrictSlash(true)
	Api = api
	
	// auth
	api.Router.HandleFunc("/signin", Chain(auth.SignIn, Method("POST"), Logging())).Methods("POST")
	api.Router.HandleFunc("/signup", Chain(auth.SignUp, Method("POST"), Logging())).Methods("POST")
	api.Router.HandleFunc("/userDetails", Chain(auth.GetUserDetails, Method("GET"), Logging()))

	// operations
	api.Router.HandleFunc("/operations", Chain(endpoints.GetOperations, Method("GET"), Logging(), Token()))
	api.Router.HandleFunc("/operation/{id}", Chain(endpoints.DeleteOperation, Method("DELETE"), Logging(), Token())).Methods("DELETE")
	api.Router.HandleFunc("/operation/{id}", Chain(endpoints.GetOperation, Method("GET"), Logging(), Token()))
	api.Router.HandleFunc("/operation", Chain(endpoints.AddOperation, Method("POST"), Logging(), Token())).Methods("POST")

	// category
	api.Router.HandleFunc("/categories", Chain(endpoints.GetAllCategories, Method("GET"), Logging(), Token()))
	api.Router.HandleFunc("/category/{id}", Chain(endpoints.GetCategory, Method("GET"), Logging(), Token()))
	api.Router.HandleFunc("/category", Chain(endpoints.AddCategory, Method("POST"), Logging(), Token())).Methods("POST")
	api.Router.HandleFunc("/mostUsedCategories", Chain(endpoints.GetMostUsedCategories, Method("GET"), Logging(), Token()))
	// router.HandleFunc("/statistic", category.GetCategoriesStatistic)

	// wallet
	api.Router.HandleFunc("/wallets", Chain(endpoints.GetWallets, Method("GET"), Logging(), Token()))
	api.Router.HandleFunc("/wallet", Chain(endpoints.AddWallet, Method("POST"), Logging(), Token())).Methods("POST")
	api.Router.HandleFunc("/wallet/{id}", Chain(endpoints.DeleteWallet, Method("DELETE"), Logging(), Token())).Methods("DELETE")
	api.Router.HandleFunc("/wallet/{id}", Chain(endpoints.GetWallet, Method("GET"), Logging(), Token()))

	log.Fatal(http.ListenAndServe(getPort(), api.Router))
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
