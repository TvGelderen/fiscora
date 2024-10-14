-- name: CreateTransaction :one
INSERT INTO transactions (user_id, amount, description, type, date)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateTransaction :exec
UPDATE transactions
SET amount = $3, description = $4, type = $5, date = $6, updated = (now() at time zone 'utc')
WHERE id = $1 AND user_id = $2;

-- name: GetTransactions :many
SELECT sqlc.embed(transactions), sqlc.embed(recurring_transactions) FROM transactions LEFT OUTER JOIN recurring_transactions ON transactions.recurring_transaction_id = recurring_transactions.id
WHERE transactions.user_id = $1
ORDER BY transactions.date
LIMIT $2
OFFSET $3;

-- name: GetTransactionsBetweenDates :many
SELECT sqlc.embed(transactions), sqlc.embed(recurring_transactions) FROM transactions LEFT OUTER JOIN recurring_transactions ON transactions.recurring_transaction_id = recurring_transactions.id
WHERE transactions.user_id = $1 AND transactions.date >= @start_date AND transactions.date <= @end_date
ORDER BY transactions.date
LIMIT $2
OFFSET $3;

-- name: GetIncomingTransactionsBetweenDates :many
SELECT sqlc.embed(transactions), sqlc.embed(recurring_transactions) FROM transactions LEFT OUTER JOIN recurring_transactions ON transactions.recurring_transaction_id = recurring_transactions.id
WHERE transactions.user_id = $1 AND amount > 0 AND transactions.date >= @start_date AND transactions.date <= @end_date
ORDER BY transactions.date
LIMIT $2
OFFSET $3;

-- name: GetOutgoingTransactionsBetweenDates :many
SELECT sqlc.embed(transactions), sqlc.embed(recurring_transactions) FROM transactions LEFT OUTER JOIN recurring_transactions ON transactions.recurring_transaction_id = recurring_transactions.id
WHERE transactions.user_id = $1 AND amount > 0 AND transactions.date >= @start_date AND transactions.date <= @end_date
ORDER BY transactions.date
LIMIT $2
OFFSET $3;

-- name: DeleteTransaction :exec
DELETE FROM transactions 
WHERE id = $1 AND user_id = $2;


-- name: CreateRecurringTransaction :one
INSERT INTO recurring_transactions (start_date, end_date, interval, days_interval)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateRecurringTransaction :exec
UPDATE recurring_transactions 
SET start_date = $3, end_date = $4, interval = $5, days_interval = $6, updated = (now() at time zone 'utc')
WHERE id = $1 AND user_id = $2;

-- name: DeleteRecurringTransaction :exec
DELETE FROM recurring_transactions 
WHERE id = $1 AND user_id = $2;
