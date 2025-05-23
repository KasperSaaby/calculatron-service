package operations

import (
	"math"

	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
)

type PowerOperation struct {
	BaseOperation
}

func NewPowerOperation() *PowerOperation {
	return &PowerOperation{
		BaseOperation: BaseOperation{
			operationType: values.OperationType_Power,
		},
	}
}

func (op *PowerOperation) Execute(precision int, operands ...float64) (float64, error) {
	if err := op.Validate(operands...); err != nil {
		return 0, err
	}

	result := math.Pow(operands[0], operands[1])
	return op.RoundResult(result, precision), nil
}

func (op *PowerOperation) Validate(operands ...float64) error {
	if len(operands) != 2 {
		return values.NewDomainError("power operation requires exactly two operands (base and exponent)", values.Code_TwoOperandsRequired)
	}
	return nil
}
