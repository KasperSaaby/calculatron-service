package operations

import (
	"math"

	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
)

type Operation interface {
	Execute(precision int, operands ...float64) (float64, error)
	Validate(operands ...float64) error
	RoundResult(value float64, precision int) float64
}

type BaseOperation struct {
	operationType values.OperationType
}

func (op *BaseOperation) Type() values.OperationType {
	return op.operationType
}

func (op *BaseOperation) RoundResult(value float64, precision int) float64 {
	multiplier := math.Pow(10, float64(precision))
	return math.Round(value*multiplier) / multiplier
}
