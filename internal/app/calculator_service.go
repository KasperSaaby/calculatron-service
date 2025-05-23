package app

import (
	"context"
	"time"

	"github.com/KasperSaaby/calculatron-service/internal/app/models"
	"github.com/KasperSaaby/calculatron-service/internal/domain/operations"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
)

type inputValidator interface {
	Validate(input models.CalculationInput) error
}

type CalculatorService struct {
	operationFactory operations.OperationFactory
	validator        inputValidator
}

func NewCalculatorService(operationFactory operations.OperationFactory, validator inputValidator) *CalculatorService {
	return &CalculatorService{
		operationFactory: operationFactory,
		validator:        validator,
	}
}

func (s *CalculatorService) PerformCalculation(_ context.Context, input models.CalculationInput) (models.CalculationResult, error) {
	err := s.validator.Validate(input)
	if err != nil {
		return models.CalculationResult{}, err
	}

	operation, err := s.operationFactory.CreateOperation(input.OperationType())
	if err != nil {
		return models.CalculationResult{}, err
	}

	result, err := operation.Execute(input.Precision(), input.Operands()...)
	if err != nil {
		return models.CalculationResult{}, err
	}

	return models.CalculationResult{
		Result:      result,
		Precision:   input.Precision(),
		OperationID: values.NewOperationID(),
		Timestamp:   time.Now(),
	}, nil
}
