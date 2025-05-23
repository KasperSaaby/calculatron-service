package operations

import "github.com/KasperSaaby/calculatron-service/internal/domain/values"

type MultiplyOperation struct {
	BaseOperation
}

func NewMultiplyOperation() *MultiplyOperation {
	return &MultiplyOperation{
		BaseOperation: BaseOperation{
			operationType: values.OperationType_Multiply,
		},
	}
}

func (op *MultiplyOperation) Execute(precision int, operands ...float64) (float64, error) {
	if err := op.Validate(operands...); err != nil {
		return 0, err
	}

	result := 1.0
	for _, operand := range operands {
		result *= operand
	}
	return op.RoundResult(result, precision), nil
}

func (op *MultiplyOperation) Validate(operands ...float64) error {
	if len(operands) < 2 {
		return values.NewDomainError("multiplication requires at least two operands", values.Code_TwoOperandsRequired)
	}
	return nil
}
