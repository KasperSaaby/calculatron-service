package operations

import "github.com/KasperSaaby/calculatron-service/internal/domain/values"

type DivideOperation struct {
	BaseOperation
}

func NewDivideOperation() *DivideOperation {
	return &DivideOperation{
		BaseOperation: BaseOperation{
			operationType: values.OperationType_Divide,
		},
	}
}

func (op *DivideOperation) Execute(precision int, operands ...float64) (float64, error) {
	if err := op.Validate(operands...); err != nil {
		return 0, err
	}

	result := operands[0] / operands[1]
	return op.RoundResult(result, precision), nil
}

func (op *DivideOperation) Validate(operands ...float64) error {
	if len(operands) != 2 {
		return values.NewDomainError("division requires exactly two operands", values.Code_TwoOperandsRequired)
	}
	if operands[1] == 0 {
		return values.NewDomainError("cannot divide by zero", values.Code_DivisionByZero)
	}
	return nil
}
