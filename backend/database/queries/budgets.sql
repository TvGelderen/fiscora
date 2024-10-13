-- name: CreateBudget :one
INSERT INTO budgets (id, user_id, name, description, amount, start_date, end_date)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateBudget :one
UPDATE budgets
SET name = $3, description = $4, amount = $5, start_date = $6, end_date = $7, updated = (now() at time zone 'utc')
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: GetBudgets :many
SELECT * FROM budgets
WHERE user_id = $1
ORDER BY created DESC
LIMIT $2
OFFSET $3;

-- name: GetBudgetsWithExpenses :many
SELECT e.id, e.budget_id, e.name, e.allocated_amount, e.current_amount FROM budgets b JOIN budget_expenses e ON b.id = e.budget_id
WHERE b.user_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteBudget :exec
DELETE FROM budgets 
WHERE id = $1 AND user_id = $2;


-- name: CreateBudgetExpense :one
INSERT INTO budget_expenses (budget_id, name, allocated_amount)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateBudgetExpense :one
UPDATE budget_expenses 
SET name = $2, allocated_amount = $3, current_amount = $4
WHERE id = $1
RETURNING *;
