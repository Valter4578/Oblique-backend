package category

import (
	"encoding/json"
	"log"
	"net/http"

	"oblique/model"

	"github.com/gorilla/mux"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllCategories")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(model.Categories)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("GetCategory")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	title := params["title"]

	category, _ := FindCategory(title)

	json.NewEncoder(w).Encode(category)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("AddCategory")

	params := r.URL.Query()

	var category model.Category

	category.Title = params.Get("title")
	category.ImageName = params.Get("imageName")
	category.Color = params.Get("color")

	model.Categories = append(model.Categories, category)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Categories)
}

func GetMostUsedCategories(w http.ResponseWriter, r *http.Request) {
	log.Println("GetMostUsedCategories")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MostUsedCategories())
}

func GetCategoriesStatistic(w http.ResponseWriter, r *http.Request) {
	log.Println("GetMostUsedCategories")

	w.Header().Set("Content-Type", "application/json")

	var categories []model.Category
	for _, category := range model.Categories {
		category.Percantage = CalculatePercantage(category)
		categories = append(categories, category)
	}
	json.NewEncoder(w).Encode(categories)
	categories = []model.Category{}
}
