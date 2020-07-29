package main

import (
	"fmt"
	"net/http"
	// "go.mongodbss.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	route()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pisya")
}
