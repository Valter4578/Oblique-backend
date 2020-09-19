package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"

	"oblique/iternal/app/db"
	"oblique/iternal/app/logger"
	"oblique/iternal/app/model"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllCategories")

	w.Header().Set("Content-Type", "application/json")

	categories, err := db.GetCategories()
	if err != nil {
		w.Write([]byte(logger.JSONError(err)))
		w.WriteHeader(500)
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
		log.Println(err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	category, err := db.GetCategory(id)
	if err != nil {
		w.Write([]byte(logger.JSONError(err)))
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(&category)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("AddCategory")

	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Println(err)
		w.Write([]byte(logger.JSONError(err)))
		return
	}

	result := db.InsertCategory(&category)
	json.NewEncoder(w).Encode(result)
}

func GetMostUsedCategories(w http.ResponseWriter, r *http.Request) {
	log.Println("GetMostUsedCategories")

	categories, err := db.GetCategories()
	if err != nil {
		w.Write([]byte(logger.JSONError(err)))
		w.WriteHeader(500)
		return
	}

	mostUsedCategories := *categories

	sort.SliceStable(mostUsedCategories, func(i, j int) bool {
		return len(mostUsedCategories[i].Operations) > len(mostUsedCategories[j].Operations)
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&mostUsedCategories)
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
