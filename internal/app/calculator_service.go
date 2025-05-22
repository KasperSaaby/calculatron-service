package app

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/KasperSaaby/calculatron-service/internal/domain/operations"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/KasperSaaby/calculatron-service/internal/store"
)

type CalculatorService struct {
	historyStore store.HistoryStore
}

func NewCalculatorService(historyStore store.HistoryStore) *CalculatorService {
	return &CalculatorService{
		historyStore: historyStore,
	}
}

func (s *CalculatorService) PerformCalculation(ctx context.Context, operationType values.OperationType, operands []float64, precision int) (values.CalculationResult, error) {
	if len(operands) == 0 {
		return values.CalculationResult{}, newAppError("no operands provided")
	}

	if precision < 0 {
		return values.CalculationResult{}, newAppError("precision cannot be negative")
	}

	op, exist := operations.Catalogue[operationType]
	if !exist {
		return values.CalculationResult{}, newAppError("operation %q is not supported", operationType)
	}

	result, err := op(operands...)
	if err != nil {
		return values.CalculationResult{}, fmt.Errorf("execute %q operation: %w", operationType, err)
	}

	roundedResult, err := s.roundFloat(result, precision)
	if err != nil {
		return values.CalculationResult{}, fmt.Errorf("round result: %w", err)
	}

	operationID := values.NewOperationID()
	err = s.historyStore.SaveCalculation(ctx, values.HistoryEntry{
		OperationID:   operationID,
		OperationType: operationType,
		Operands:      operands,
		Result:        roundedResult,
		Precision:     int32(precision),
		Timestamp:     time.Now(),
	})
	if err != nil {
		return values.CalculationResult{}, fmt.Errorf("save to calculation history: %w", err)
	}

	return values.CalculationResult{
		Result:      roundedResult,
		Precision:   precision,
		OperationID: operationID,
		Timestamp:   time.Now(),
	}, nil
}

func (s *CalculatorService) roundFloat(val float64, precision int) (float64, error) {
	multiplier := math.Pow(10, float64(precision))
	rounded := math.Round(val*multiplier) / multiplier
	return rounded, nil
}
