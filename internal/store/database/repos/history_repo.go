package repos

import (
	"context"
	"database/sql"
	"time"

	db "github.com/KasperSaaby/calculatron-service/generated/database/history"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/sqlc-dev/pqtype"
)

type HistoryRepo struct {
	querier *db.Queries
}

func NewHistoryRepo(conn *sql.DB) *HistoryRepo {
	return &HistoryRepo{
		querier: db.New(conn),
	}
}

func (r *HistoryRepo) SaveCalculation(
	ctx context.Context,
	operationID values.OperationID,
	operationType values.OperationType,
	operands []float64,
	result float64,
	precision int,
) error {
	return r.querier.Insert(ctx, db.InsertParams{
		OperationID:   operationID.String(),
		OperationType: operationType.String(),
		Operands:      operands,
		Result:        result,
		Precision:     int32(precision),
		Timestamp:     time.Now(),
		Metadata:      pqtype.NullRawMessage{},
	})
}
