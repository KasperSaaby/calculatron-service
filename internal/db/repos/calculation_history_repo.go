package repos

import (
	"context"
	"database/sql"
	"time"

	"github.com/KasperSaaby/calculatron-service/generated/database/calculation_history"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/sqlc-dev/pqtype"
)

type CalculationHistoryRepo struct {
	querier *calculation_history.Queries
}

func NewCalculationHistoryRepo(conn *sql.DB) *CalculationHistoryRepo {
	return &CalculationHistoryRepo{
		querier: calculation_history.New(conn),
	}
}

func (r *CalculationHistoryRepo) SaveCalculation(
	ctx context.Context,
	operationID values.OperationID,
	operationType values.OperationType,
	operands []float64,
	result float64,
	precision int,
) error {
	return r.querier.Insert(ctx, calculation_history.InsertParams{
		OperationID:   operationID.String(),
		OperationType: operationType.String(),
		Operands:      operands,
		Result:        result,
		Precision:     int32(precision),
		Timestamp:     time.Now(),
		Metadata:      pqtype.NullRawMessage{},
	})
}
