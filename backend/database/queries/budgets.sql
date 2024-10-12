-- name: CreateBudget :one
INSERT INTO budgets (id, user_id, name, description, amount, start_date, end_date, created, updated)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: UpdateBudget :exec
UPDATE budgets
SET name = $3, description = $4, amount = $5, start_date = $6, end_date = $7, updated = $8
WHERE id = $1 AND user_id = $2;

-- name: GetBudgets :many
SELECT * FROM budgets JOIN budget_expenses ON budgets.id = budget_expenses.budget_id
WHERE budgets.user_id = $1
LIMIT $2
OFFSET $3;


-- name: CreateBudgetExpense :one
INSERT INTO budget_expenses (budget_id, name, description, allocated_amount, current_amount)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateBudgetExpense :one
UPDATE budget_expenses 
SET name = $2, description = $3, allocated_amount = $4, current_amount = $5
WHERE id = $1
RETURNING *;
