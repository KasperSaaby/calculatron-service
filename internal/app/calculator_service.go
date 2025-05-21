package app

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/KasperSaaby/calculatron-service/internal/domain/operations"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/KasperSaaby/calculatron-service/internal/store/database/repos"
	"github.com/google/uuid"
)

type CalculatorService struct {
	historyRepo *repos.HistoryRepo
}

func NewCalculatorService(historyRepo *repos.HistoryRepo) *CalculatorService {
	return &CalculatorService{
		historyRepo: historyRepo,
	}
}

func (s *CalculatorService) PerformCalculation(ctx context.Context, operationType values.OperationType, operands []float64, precision int) (Result, error) {
	if len(operands) == 0 {
		return Result{}, newAppError("no operands provided")
	}

	if precision < 0 {
		return Result{}, newAppError("precision cannot be negative")
	}

	op, exist := operations.Catalogue[operationType]
	if !exist {
		return Result{}, newAppError("operation %q is not supported", operationType)
	}

	result, err := op(operands...)
	if err != nil {
		return Result{}, fmt.Errorf("execute %q operation: %w", operationType, err)
	}

	roundedResult, err := s.roundFloat(result, precision)
	if err != nil {
		return Result{}, fmt.Errorf("round result: %w", err)
	}

	err = s.historyRepo.SaveCalculation(ctx, values.NewOperationID(), operationType, operands, roundedResult, precision)
	if err != nil {
		return Result{}, fmt.Errorf("save to calculation history: %w", err)
	}

	return Result{
		Result:      roundedResult,
		Precision:   precision,
		OperationID: uuid.NewString(),
		Timestamp:   time.Now(),
	}, nil
}

func (s *CalculatorService) roundFloat(val float64, precision int) (float64, error) {
	multiplier := math.Pow(10, float64(precision))
	rounded := math.Round(val*multiplier) / multiplier
	return rounded, nil
}
