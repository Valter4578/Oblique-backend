package category

import (
	"mellow/model"
	"sort"
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

// MostUsedCategories is method that returns slice of most used categories based on number of expenses
func MostUsedCategories() []model.Category {
	mostUsedCategories := model.Categories

	sort.SliceStable(mostUsedCategories, func(i, j int) bool {
		return len(mostUsedCategories[i].Expenses) > len(mostUsedCategories[j].Expenses)
	})

	return mostUsedCategories
}
