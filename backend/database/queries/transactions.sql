-- name: CreateTransaction :one
INSERT INTO transactions (user_id, amount, description, type, date)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateTransaction :exec
UPDATE transactions
SET amount = $3, description = $4, type = $5, date = $6, updated = (now() at time zone 'utc')
WHERE id = $1 AND user_id = $2;

-- name: GetTransactions :many
SELECT * FROM transactions t LEFT OUTER JOIN recurring_transactions rt ON t.recurring_transaction_id = rt.id
WHERE t.user_id = $1
ORDER BY t.date
LIMIT $2
OFFSET $3;

-- name: GetTransactionsBetweenDates :many
SELECT * FROM transactions t LEFT OUTER JOIN recurring_transactions rt ON t.recurring_transaction_id = rt.id
WHERE t.user_id = $1 AND t.date >= $2 AND t.date <= $3
ORDER BY t.date
LIMIT $4
OFFSET $5;

-- name: GetIncomingTransactionsBetweenDates :many
SELECT * FROM transactions t LEFT OUTER JOIN recurring_transactions rt ON t.recurring_transaction_id = rt.id
WHERE t.user_id = $1 AND t.amount > 0 AND t.date >= $2 AND t.date <= $3
ORDER BY t.date
LIMIT $4
OFFSET $5;

-- name: GetOutgoingTransactionsBetweenDates :many
SELECT * FROM transactions t LEFT OUTER JOIN recurring_transactions rt ON t.recurring_transaction_id = rt.id
WHERE t.user_id = $1 AND t.amount < 0 AND t.date >= $2 AND t.date <= $3
ORDER BY t.date
LIMIT $4
OFFSET $5;

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
