package operations

import "github.com/KasperSaaby/calculatron-service/internal/domain/values"

type OperationFactory interface {
	CreateOperation(operationType values.OperationType) (Operation, error)
}

type operationFactory struct{}

func NewOperationFactory() OperationFactory {
	return &operationFactory{}
}

func (f *operationFactory) CreateOperation(operationType values.OperationType) (Operation, error) {
	switch operationType {
	case values.OperationType_Add:
		return NewAddOperation(), nil
	case values.OperationType_Subtract:
		return NewSubtractOperation(), nil
	case values.OperationType_Multiply:
		return NewMultiplyOperation(), nil
	case values.OperationType_Divide:
		return NewDivideOperation(), nil
	case values.OperationType_Power:
		return NewPowerOperation(), nil
	default:
		return nil, values.NewDomainError("unsupported operation type", values.Code_UnsupportedOperation)
	}
}
