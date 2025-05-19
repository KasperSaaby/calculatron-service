package calculator

import (
	"calculatron/internal/db/repos"
	"calculatron/internal/domain/operations"
	"calculatron/internal/domain/values"
	"context"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	calculationHistoryRepo *repos.CalculationHistoryRepo
}

func NewService(calculationHistoryRepo *repos.CalculationHistoryRepo) *Service {
	return &Service{
		calculationHistoryRepo: calculationHistoryRepo,
	}
}

func (s *Service) PerformCalculation(ctx context.Context, operationType values.OperationType, operands []float64, precision int) (Result, error) {
	if len(operands) == 0 {
		return Result{}, newClientError(fmt.Errorf("no operands provided"))
	}

	if precision < 0 {
		return Result{}, newClientError(fmt.Errorf("precision cannot be negative"))
	}

	op, exist := operations.Catalogue[operationType]
	if !exist {
		return Result{}, newClientError(fmt.Errorf("operation %q is not supported", operationType))
	}

	result, err := op(operands...)
	if err != nil {
		return Result{}, fmt.Errorf("execute %q operation: %w", operationType, err)
	}

	roundedResult, err := s.roundFloat(result, precision)
	if err != nil {
		return Result{}, fmt.Errorf("round result: %w", err)
	}

	err = s.calculationHistoryRepo.SaveCalculation(ctx, values.NewOperationID(), operationType, operands, roundedResult, precision)
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

func (s *Service) roundFloat(val float64, precision int) (float64, error) {
	multiplier := math.Pow(10, float64(precision))
	rounded := math.Round(val*multiplier) / multiplier
	return rounded, nil
}
