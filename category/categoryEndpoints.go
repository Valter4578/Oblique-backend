package category

import (
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"

	"oblique/db"
	"oblique/logger"
	"oblique/model"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllCategories")

	w.Header().Set("Content-Type", "application/json")

	var categories []model.Category

	err := db.GetCategories(&categories)
	if err != nil {
		logger.LogError(&err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("GetCategory")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		logger.LogError(&err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	var category model.Category
	err = db.GetCategory(id, &category)
	if err != nil {
		logger.LogError(&err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	json.NewEncoder(w).Encode(&category)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("AddCategory")

	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		logger.LogError(&err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	result := db.InsertCategory(&category)
	json.NewEncoder(w).Encode(result)
}

func GetMostUsedCategories(w http.ResponseWriter, r *http.Request) {
	log.Println("GetMostUsedCategories")

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(MostUsedCategories())
}

// func GetCategoriesStatistic(w http.ResponseWriter, r *http.Request) {
// 	log.Println("GetMostUsedCategories")

// 	w.Header().Set("Content-Type", "application/json")

// 	var categories []model.Category
// 	for _, category := range model.Categories {
// 		category.Percantage = CalculatePercantage(category)
// 		categories = append(categories, category)
// 	}
// 	json.NewEncoder(w).Encode(categories)
// 	categories = []model.Category{}
// }
