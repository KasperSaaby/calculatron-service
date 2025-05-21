package repository

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

func (r *HistoryRepo) Create(ctx context.Context, entry values.HistoryEntry) error {
	return r.querier.Insert(ctx, db.InsertParams{
		OperationID:   entry.OperationID.String(),
		OperationType: entry.OperationType.String(),
		Operands:      entry.Operands,
		Result:        entry.Result,
		Precision:     entry.Precision,
		Timestamp:     time.Now(),
		Metadata:      pqtype.NullRawMessage{},
	})
}

func (r *HistoryRepo) FindAll(ctx context.Context, offset, limit int) ([]values.HistoryEntry, error) {
	entities, err := r.querier.SelectAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []values.HistoryEntry
	for _, entity := range entities {
		result = append(result, r.mapToDomain(entity))
	}

	return result, nil
}

func (*HistoryRepo) mapToDomain(history db.History) values.HistoryEntry {
	return values.HistoryEntry{
		OperationID:   values.OperationID(history.OperationID),
		OperationType: values.OperationType(history.OperationType),
		Operands:      history.Operands,
		Result:        history.Result,
		Precision:     history.Precision,
		Timestamp:     history.Timestamp,
	}
}
