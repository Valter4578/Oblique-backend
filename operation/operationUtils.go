package expense

import (
	"mellow/model"
)

func allOperations() []model.Operation {
	var operations []model.Operation
	for _, category := range model.Categories {

		for _, operation := range category.Operations {
			operations = append(operations, operation)
		}
	}

	return operations
}
