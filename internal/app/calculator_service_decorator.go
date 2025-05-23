package app

import (
	"context"
	"time"

	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/KasperSaaby/calculatron-service/internal/store"
)

type Calculator interface {
	PerformCalculation(ctx context.Context, operationType values.OperationType, operands []float64, precision int) (values.CalculationResult, error)
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

func (d *CalculatorServiceDecorator) PerformCalculation(ctx context.Context, operationType values.OperationType, operands []float64, precision int) (values.CalculationResult, error) {
	result, err := d.calculator.PerformCalculation(ctx, operationType, operands, precision)
	if err != nil {
		return values.CalculationResult{}, err
	}

	err = d.store.SaveCalculation(ctx, values.HistoryEntry{
		OperationID:   result.OperationID,
		OperationType: operationType,
		Operands:      operands,
		Result:        result.Result,
		Precision:     precision,
		Timestamp:     time.Now(),
	})
	if err != nil {
		return values.CalculationResult{}, err
	}

	return result, nil
}
