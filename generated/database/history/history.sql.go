// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: history.sql

package history

import (
	"context"
	"time"

	"github.com/lib/pq"
	"github.com/sqlc-dev/pqtype"
)

const findAll = `-- name: FindAll :many
SELECT operation_id, operation_type, operands, result, precision, timestamp, metadata FROM history LIMIT $1 OFFSET $2
`

type FindAllParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) FindAll(ctx context.Context, arg FindAllParams) ([]History, error) {
	rows, err := q.db.QueryContext(ctx, findAll, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []History
	for rows.Next() {
		var i History
		if err := rows.Scan(
			&i.OperationID,
			&i.OperationType,
			pq.Array(&i.Operands),
			&i.Result,
			&i.Precision,
			&i.Timestamp,
			&i.Metadata,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findByID = `-- name: FindByID :one
SELECT operation_id, operation_type, operands, result, precision, timestamp, metadata FROM history WHERE operation_id = $1
`

func (q *Queries) FindByID(ctx context.Context, operationID string) (History, error) {
	row := q.db.QueryRowContext(ctx, findByID, operationID)
	var i History
	err := row.Scan(
		&i.OperationID,
		&i.OperationType,
		pq.Array(&i.Operands),
		&i.Result,
		&i.Precision,
		&i.Timestamp,
		&i.Metadata,
	)
	return i, err
}

const insert = `-- name: Insert :exec
INSERT INTO history
    (operation_id, operation_type, operands, result, precision, timestamp, metadata)
VALUES
    ($1, $2, $3,$4,$5,$6,$7)
`

type InsertParams struct {
	OperationID   string
	OperationType string
	Operands      []float64
	Result        float64
	Precision     int32
	Timestamp     time.Time
	Metadata      pqtype.NullRawMessage
}

func (q *Queries) Insert(ctx context.Context, arg InsertParams) error {
	_, err := q.db.ExecContext(ctx, insert,
		arg.OperationID,
		arg.OperationType,
		pq.Array(arg.Operands),
		arg.Result,
		arg.Precision,
		arg.Timestamp,
		arg.Metadata,
	)
	return err
}
