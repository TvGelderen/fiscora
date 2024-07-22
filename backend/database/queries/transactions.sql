-- name: CreateTransaction :one
INSERT INTO transactions (user_id, amount, description, incoming, type, recurring, start_date, end_date, interval, days_interval, created, updated)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: UpdateTransaction :exec
UPDATE transactions
SET amount = $3, description = $4, incoming = $5, type = $6, recurring = $7, start_date = $8, end_date = $9, interval = $10, days_interval = $11, updated = $12
WHERE id = $1 AND user_id = $2;

-- name: GetUserTransactions :many
SELECT * FROM transactions 
WHERE user_id = $1;

-- name: GetUserTransactionsBetweenDates :many
SELECT * FROM transactions 
WHERE user_id = $1 AND start_date <= $2 AND end_date >= $3;

-- name: GetUserIncomingTransactions :many
SELECT * FROM transactions 
WHERE user_id = $1 AND incoming = 1;

-- name: GetUserOutgoingTransactions :many
SELECT * FROM transactions 
WHERE user_id = $1 AND incoming = 0;

-- name: GetUserTransactionsByType :many
SELECT * FROM transactions 
WHERE user_id = $1 AND type = $2;

-- name: DeleteTransaction :exec
DELETE FROM transactions 
WHERE id = $1 AND user_id = $2;
