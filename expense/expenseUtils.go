package expense

import (
	"mellow/model"
)

func allExpenses() []model.Expense {
	var expenses []model.Expense
	for _, category := range model.Categories {

		for _, expense := range category.Expenses {
			expenses = append(expenses, expense)
		}
	}

	return expenses
}
