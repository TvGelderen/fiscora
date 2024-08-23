-- name: CreateTransaction :one
INSERT INTO transactions (user_id, amount, description, type, recurring, start_date, end_date, interval, days_interval, created, updated)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: UpdateTransaction :exec
UPDATE transactions
SET amount = $3, description = $4, type = $5, recurring = $6, start_date = $7, end_date = $8, interval = $9, days_interval = $10, updated = $11
WHERE id = $1 AND user_id = $2;

-- name: GetUserTransactions :many
SELECT * FROM transactions 
WHERE user_id = $1
ORDER BY start_date
LIMIT $2
OFFSET $3;

-- name: GetUserTransactionsBetweenDates :many
SELECT * FROM transactions 
WHERE user_id = $1 AND start_date <= $2 AND end_date >= $3
ORDER BY start_date
LIMIT $4
OFFSET $5;

-- name: GetUserIncomeTransactionsBetweenDates :many
SELECT * FROM transactions 
WHERE user_id = $1 AND amount > 0 AND start_date <= $2 AND end_date >= $3
ORDER BY start_date
LIMIT $4
OFFSET $5;

-- name: GetUserExpenseTransactionsBetweenDates :many
SELECT * FROM transactions 
WHERE user_id = $1 AND amount < 0 AND start_date <= $2 AND end_date >= $3
ORDER BY start_date
LIMIT $4
OFFSET $5;

-- name: GetUserIncomingTransactions :many
SELECT * FROM transactions 
WHERE user_id = $1 AND amount > 0;

-- name: GetUserOutgoingTransactions :many
SELECT * FROM transactions 
WHERE user_id = $1 AND amount < 0;

-- name: GetUserTransactionsByType :many
SELECT * FROM transactions 
WHERE user_id = $1 AND type = $2;

-- name: DeleteTransaction :exec
DELETE FROM transactions 
WHERE id = $1 AND user_id = $2;
