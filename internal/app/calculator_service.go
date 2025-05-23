package app

import (
	"context"
	"fmt"
	"time"

	"github.com/KasperSaaby/calculatron-service/internal/domain/operations"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
)

type CalculatorService struct {
	operationFactory operations.OperationFactory
}

func NewCalculatorService(operationFactory operations.OperationFactory) *CalculatorService {
	return &CalculatorService{
		operationFactory: operationFactory,
	}
}

func (s *CalculatorService) PerformCalculation(_ context.Context, operationType values.OperationType, operands []float64, precision int) (values.CalculationResult, error) {
	if len(operands) == 0 {
		return values.CalculationResult{}, newAppError("no operands provided")
	}

	if precision < 0 {
		return values.CalculationResult{}, newAppError("precision cannot be negative")
	}

	operation, err := s.operationFactory.CreateOperation(operationType)
	if err != nil {
		return values.CalculationResult{}, err
	}

	result, err := operation.Execute(precision, operands...)
	if err != nil {
		return values.CalculationResult{}, fmt.Errorf("execute %q operation: %w", operationType, err)
	}

	return values.CalculationResult{
		Result:      result,
		Precision:   precision,
		OperationID: values.NewOperationID(),
		Timestamp:   time.Now(),
	}, nil
}
