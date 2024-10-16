-- name: CreateTransaction :one
INSERT INTO transactions (user_id, recurring_transaction_id, amount, description, type, date)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateTransaction :exec
UPDATE transactions
SET amount = $3, description = $4, type = $5, date = $6, updated = (now() at time zone 'utc')
WHERE id = $1 AND user_id = $2;

-- name: GetBaseTransactionsBetweenDates :many
SELECT * FROM transactions
WHERE user_id = $1 AND date >= sqlc.arg(start_date) AND date <= sqlc.arg(end_date)
ORDER BY date
LIMIT $2
OFFSET $3;

-- name: GetTransactionsBetweenDates :many
SELECT sqlc.embed(ft) FROM full_transaction ft
WHERE ft.user_id = $1 AND ft.date >= sqlc.arg(start_date) AND ft.date <= sqlc.arg(end_date)
ORDER BY ft.date
LIMIT $2
OFFSET $3;

-- name: GetIncomeTransactionsBetweenDates :many
SELECT sqlc.embed(ft) FROM full_transaction ft
WHERE ft.user_id = $1 AND ft.amount > 0 AND ft.date >= sqlc.arg(start_date) AND ft.date <= sqlc.arg(end_date)
ORDER BY ft.date
LIMIT $2
OFFSET $3;

-- name: GetExpenseTransactionsBetweenDates :many
SELECT sqlc.embed(ft) FROM full_transaction ft
WHERE ft.user_id = $1 AND ft.amount < 0 AND ft.date >= sqlc.arg(start_date) AND ft.date <= sqlc.arg(end_date)
ORDER BY ft.date
LIMIT $2
OFFSET $3;

-- name: DeleteTransaction :exec
DELETE FROM transactions 
WHERE id = $1 AND user_id = $2;


-- name: CreateRecurringTransaction :one
INSERT INTO recurring_transactions (user_id, start_date, end_date, interval, days_interval)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateRecurringTransaction :exec
UPDATE recurring_transactions 
SET start_date = $3, end_date = $4, interval = $5, days_interval = $6, updated = (now() at time zone 'utc')
WHERE id = $1 AND user_id = $2;

-- name: DeleteRecurringTransaction :exec
DELETE FROM recurring_transactions 
WHERE id = $1 AND user_id = $2;
