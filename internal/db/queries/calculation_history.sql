-- name: Insert :exec
INSERT INTO calculation_history
    (operation_id, operation_type, operands, result, precision, timestamp, metadata)
VALUES
    ($1, $2, $3,$4,$5,$6,$7);

-- name: QueryAll :many
SELECT * FROM calculation_history;
