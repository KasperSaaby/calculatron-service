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
	default:
		return nil, values.NewDomainError("unsupported operation type", values.Code_UnsupportedOperation)
	}
}
