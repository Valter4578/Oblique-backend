package category

import (
	"mellow/model"
	"strings"
)

// FindCategory finds a category by its title. Returns category and index
func FindCategory(title string) (model.Category, int) {
	var category model.Category
	var index int

	for i, c := range model.Categories {
		if strings.ToLower(title) == strings.ToLower(c.Title) {
			category = c
			index = i
		}
	}

	return category, index
}
