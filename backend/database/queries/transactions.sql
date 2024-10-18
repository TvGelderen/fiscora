-- name: CreateTransaction :one
INSERT INTO transactions (user_id, recurring_transaction_id, amount, description, type, date)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateTransaction :exec
UPDATE transactions
SET amount = $3, description = $4, type = $5, date = $6, updated = (now() at time zone 'utc')
WHERE id = $1 AND user_id = $2;

-- name: GetTransactionById :one
SELECT * FROM transactions
WHERE id = $1 AND user_id = $2
LIMIT 1;

-- name: GetTransactionsByRecurringTransactionId :many
SELECT * FROM transactions
WHERE recurring_transaction_id = sqlc.arg(recurring_transaction_id)::int AND user_id = $1;

-- name: GetTransactionAmountsBetweenDates :many
SELECT amount FROM transactions
WHERE user_id = $1 AND date >= sqlc.arg(start_date) AND date <= sqlc.arg(end_date);

-- name: GetIncomeTransactionAmountsBetweenDates :many
SELECT amount, type FROM transactions
WHERE user_id = $1 AND date >= sqlc.arg(start_date) AND date <= sqlc.arg(end_date);

-- name: GetExpenseTransactionAmountsBetweenDates :many
SELECT amount, type FROM transactions
WHERE user_id = $1 AND amount > 0 AND date >= sqlc.arg(start_date) AND date <= sqlc.arg(end_date);

-- name: GetBaseTransactionsBetweenDates :many
SELECT * FROM transactions
WHERE user_id = $1 AND amount < 0 AND date >= sqlc.arg(start_date) AND date <= sqlc.arg(end_date)
ORDER BY date
LIMIT $2
OFFSET $3;

-- name: GetTransactionsBetweenDates :many
SELECT sqlc.embed(t) FROM full_transaction t
WHERE t.user_id = $1 AND t.date >= sqlc.arg(start_date) AND t.date <= sqlc.arg(end_date)
ORDER BY t.date
LIMIT $2
OFFSET $3;

-- name: GetIncomeTransactionsBetweenDates :many
SELECT sqlc.embed(t) FROM full_transaction t
WHERE t.user_id = $1 AND t.amount > 0 AND t.date >= sqlc.arg(start_date) AND t.date <= sqlc.arg(end_date)
ORDER BY t.date
LIMIT $2
OFFSET $3;

-- name: GetExpenseTransactionsBetweenDates :many
SELECT sqlc.embed(t) FROM full_transaction t
WHERE t.user_id = $1 AND t.amount < 0 AND t.date >= sqlc.arg(start_date) AND t.date <= sqlc.arg(end_date)
ORDER BY t.date
LIMIT $2
OFFSET $3;

-- name: DeleteTransaction :exec
DELETE FROM transactions 
WHERE id = $1 AND user_id = $2;

-- name: DeleteTransactionsByRecurringTransactionId :execrows
DELETE FROM transactions 
WHERE recurring_transaction_id = sqlc.arg(recurring_transaction_id)::int AND user_id = $1;

-- name: DeleteTransactionsByRecurringTransactionIdAndWhereDate :execrows
DELETE FROM transactions 
WHERE recurring_transaction_id = sqlc.arg(recurring_transaction_id)::int AND user_id = $1 AND date > $2;


-- name: CreateRecurringTransaction :one
INSERT INTO recurring_transactions (user_id, start_date, end_date, interval, days_interval)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateRecurringTransaction :exec
UPDATE recurring_transactions 
SET start_date = $3, end_date = $4, interval = $5, days_interval = $6, updated = (now() at time zone 'utc')
WHERE id = $1 AND user_id = $2;

-- name: GetRecurringTransactionById :one
SELECT * FROM recurring_transactions
WHERE id = $1 AND user_id = $2
LIMIT 1;

-- name: DeleteRecurringTransaction :exec
DELETE FROM recurring_transactions 
WHERE id = $1 AND user_id = $2;
