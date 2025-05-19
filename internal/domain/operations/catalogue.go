package operations

import "calculatron/internal/domain/values"

type operationFunc func(...float64) (float64, error)

var Catalogue = map[values.OperationType]operationFunc{
	values.OperationType_Add:      Add,
	values.OperationType_Subtract: Subtract,
}
