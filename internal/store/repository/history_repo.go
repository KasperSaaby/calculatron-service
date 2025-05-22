package repository

import (
	"context"
	"database/sql"
	"fmt"
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
	entities, err := r.querier.FindAll(ctx, db.FindAllParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, fmt.Errorf("find history entries with offset %d and limit %d: %w", offset, limit, err)
	}

	var result []values.HistoryEntry
	for _, entity := range entities {
		result = append(result, r.mapToDomain(entity))
	}

	return result, nil
}

func (r *HistoryRepo) FindByID(ctx context.Context, operationID values.OperationID) (values.HistoryEntry, error) {
	entity, err := r.querier.FindByID(ctx, operationID.String())
	if err != nil {
		return values.HistoryEntry{}, fmt.Errorf("find history entry by id %s: %w", operationID, err)
	}

	return r.mapToDomain(entity), nil
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
