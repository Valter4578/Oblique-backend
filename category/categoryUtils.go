package category

import (
	"oblique/model"
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
		return len(mostUsedCategories[i].Operations) > len(mostUsedCategories[j].Operations)
	})

	return mostUsedCategories
}

// CalculatePercantage get the Category as parameter and return float32 value which is percantage of category compared to other categories
func CalculatePercantage(category model.Category) float32 {
	// get all operations' amount
	var totalAmount float32 // summary of all operations' amount
	for _, c := range model.Categories {
		for _, operation := range c.Operations {
			totalAmount += float32(operation.Amount)
		}
	}

	// get category's operations' amount
	var categoryAmount float32
	for _, operation := range category.Operations {
		categoryAmount += float32(operation.Amount)
	}

	// calculate and return percantage
	return (categoryAmount / totalAmount) * 100
}
