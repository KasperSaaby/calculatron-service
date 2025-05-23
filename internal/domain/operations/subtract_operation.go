package operations

import "github.com/KasperSaaby/calculatron-service/internal/domain/values"

type SubtractOperation struct {
	BaseOperation
}

func NewSubtractOperation() *SubtractOperation {
	return &SubtractOperation{
		BaseOperation: BaseOperation{
			operationType: values.OperationType_Subtract,
		},
	}
}

func (op *SubtractOperation) Execute(precision int, operands ...float64) (float64, error) {
	if err := op.Validate(operands...); err != nil {
		return 0, err
	}

	result := operands[0] - operands[1]
	return op.RoundResult(result, precision), nil
}

func (op *SubtractOperation) Validate(operands ...float64) error {
	if len(operands) != 2 {
		return values.NewDomainError("subtraction requires exactly two operands", values.Code_TwoOperandsRequired)
	}
	return nil
}
