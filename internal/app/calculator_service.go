package app

import (
	"context"
	"fmt"
	"time"

	"github.com/KasperSaaby/calculatron-service/internal/app/models"
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

func (s *CalculatorService) PerformCalculation(_ context.Context, input models.CalculationInput) (models.CalculationResult, error) {
	operation, err := s.operationFactory.CreateOperation(input.OperationType())
	if err != nil {
		return models.CalculationResult{}, err
	}

	result, err := operation.Execute(input.Precision(), input.Operands()...)
	if err != nil {
		return models.CalculationResult{}, fmt.Errorf("execute %q operation: %w", input.OperationType(), err)
	}

	return models.CalculationResult{
		Result:      result,
		Precision:   input.Precision(),
		OperationID: values.NewOperationID(),
		Timestamp:   time.Now(),
	}, nil
}
