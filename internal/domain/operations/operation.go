package operations

import "github.com/KasperSaaby/calculatron-service/internal/domain/values"

type Operation interface {
	Execute(operands ...float64) (float64, error)
	Validate(operands ...float64) error
}

type BaseOperation struct {
	operationType values.OperationType
}

func (op BaseOperation) Type() values.OperationType {
	return op.operationType
}
