package app

import (
	"context"
	"time"

	"github.com/KasperSaaby/calculatron-service/internal/app/models"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/KasperSaaby/calculatron-service/internal/store"
)

type Calculator interface {
	PerformCalculation(ctx context.Context, input models.CalculationInput) (models.CalculationResult, error)
}

type CalculatorServiceDecorator struct {
	calculator Calculator
	store      store.HistoryStore
}

func NewCalculatorServiceDecorator(calculator Calculator, store store.HistoryStore) *CalculatorServiceDecorator {
	return &CalculatorServiceDecorator{
		calculator: calculator,
		store:      store,
	}
}

func (d *CalculatorServiceDecorator) PerformCalculation(ctx context.Context, input models.CalculationInput) (models.CalculationResult, error) {
	result, err := d.calculator.PerformCalculation(ctx, input)
	if err != nil {
		return models.CalculationResult{}, err
	}

	err = d.store.SaveCalculation(ctx, values.HistoryEntry{
		OperationID:   result.OperationID,
		OperationType: input.OperationType(),
		Operands:      input.Operands(),
		Result:        result.Result,
		Precision:     input.Precision(),
		Timestamp:     time.Now(),
	})
	if err != nil {
		return models.CalculationResult{}, err
	}

	return result, nil
}
