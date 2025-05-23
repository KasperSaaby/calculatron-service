-- name: Insert :exec
INSERT INTO history
    (operation_id, operation_type, operands, result, precision, timestamp, metadata)
VALUES
    ($1, $2, $3,$4,$5,$6,$7);

-- name: FindAll :many
SELECT * FROM history LIMIT $1 OFFSET $2;

-- name: FindByID :one
SELECT * FROM history WHERE operation_id = $1;
