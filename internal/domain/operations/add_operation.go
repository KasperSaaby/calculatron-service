package operations

import "github.com/KasperSaaby/calculatron-service/internal/domain/values"

type AddOperation struct {
	BaseOperation
}

func NewAddOperation() *AddOperation {
	return &AddOperation{
		BaseOperation: BaseOperation{
			operationType: values.OperationType_Add,
		},
	}
}

func (op *AddOperation) Execute(precision int, operands ...float64) (float64, error) {
	if err := op.Validate(operands...); err != nil {
		return 0, err
	}

	result := 0.0
	for _, operand := range operands {
		result += operand
	}

	return op.RoundResult(result, precision), nil
}

func (op *AddOperation) Validate(operands ...float64) error {
	if len(operands) < 2 {
		return values.NewDomainError("addition requires at least two operands", values.Code_TwoOrMoreOperandsRequired)
	}

	return nil
}
