package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type IBudgetRepository interface {
	Get(ctx context.Context, userId uuid.UUID) (*[]BudgetWithExpenses, error)
	GetById(ctx context.Context, id string) (*BudgetWithExpenses, error)
	Add(ctx context.Context, params CreateBudgetParams) (*Budget, error)
	Update(ctx context.Context, params UpdateBudgetParams) error
	Remove(ctx context.Context, id string, userId uuid.UUID) error

	GetExpenses(ctx context.Context, budgetId string) (*[]BudgetExpense, error)
	AddExpense(ctx context.Context, params CreateBudgetExpenseParams) (*BudgetExpense, error)
	UpdateExpense(ctx context.Context, params UpdateBudgetExpenseParams) error
	RemoveExpense(ctx context.Context, id int32, budgetId string) error
}

type BudgetRepository struct {
	db *sql.DB
}

func CreateBudgetRepository(db *sql.DB) *BudgetRepository {
	return &BudgetRepository{
		db: db,
	}
}

func (repository *BudgetRepository) Get(ctx context.Context, userId uuid.UUID) (*[]BudgetWithExpenses, error) {
	db := New(repository.db)
	budgets, err := db.GetBudgets(ctx, GetBudgetsParams{
		UserID: userId,
		Limit:  MaxFetchLimit,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}
	budgetExpenses, err := db.GetBudgetsExpenses(ctx, GetBudgetsExpensesParams{
		UserID: userId,
		Limit:  MaxFetchLimit,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}

	budgetMap := make(map[string][]BudgetExpense, len(budgets))
	for _, budgetExpense := range budgetExpenses {
		budgetMap[budgetExpense.BudgetExpense.BudgetID] = append(budgetMap[budgetExpense.BudgetExpense.BudgetID], budgetExpense.BudgetExpense)
	}

	budgetsWithExpenses := make([]BudgetWithExpenses, len(budgets))
	for idx, budget := range budgets {
		budgetsWithExpenses[idx] = BudgetWithExpenses{
			Budget:   budget,
			Expenses: budgetMap[budget.ID],
		}
	}

	return &budgetsWithExpenses, nil
}

func (repository *BudgetRepository) GetById(ctx context.Context, id string) (*BudgetWithExpenses, error) {
	db := New(repository.db)
	budget, err := db.GetBudget(ctx, id)
	if err != nil {
		return nil, err
	}
	expenses, err := db.GetBudgetExpenses(ctx, id)
	if err != nil {
		return nil, err
	}

	budgetWithExpenses := BudgetWithExpenses{
		Budget:   budget,
		Expenses: expenses,
	}

	return &budgetWithExpenses, nil
}

func (repository *BudgetRepository) Add(ctx context.Context, params CreateBudgetParams) (*Budget, error) {
	db := New(repository.db)
	budget, err := db.CreateBudget(ctx, params)
	return &budget, err
}

func (repository *BudgetRepository) Update(ctx context.Context, params UpdateBudgetParams) error {
	db := New(repository.db)
	// TODO: Handle updating startDate and endDate
	_, err := db.UpdateBudget(ctx, params)
	return err
}

func (repository *BudgetRepository) Remove(ctx context.Context, id string, userId uuid.UUID) error {
	db := New(repository.db)
	err := db.DeleteBudget(ctx, DeleteBudgetParams{
		ID:     id,
		UserID: userId,
	})
	return err
}

func (repository *BudgetRepository) GetExpenses(ctx context.Context, budgetId string) (*[]BudgetExpense, error) {
	db := New(repository.db)
	expenses, err := db.GetBudgetExpenses(ctx, budgetId)
	if err != nil {
		return nil, err
	}

	return &expenses, nil
}

func (repository *BudgetRepository) AddExpense(ctx context.Context, params CreateBudgetExpenseParams) (*BudgetExpense, error) {
	db := New(repository.db)
	budgetExpense, err := db.CreateBudgetExpense(ctx, params)
	return &budgetExpense, err
}

func (repository *BudgetRepository) UpdateExpense(ctx context.Context, params UpdateBudgetExpenseParams) error {
	db := New(repository.db)
	_, err := db.UpdateBudgetExpense(ctx, params)
	return err
}

func (repository *BudgetRepository) RemoveExpense(ctx context.Context, id int32, budgetId string) error {
	db := New(repository.db)
	return db.DeleteBudgetExpense(ctx, DeleteBudgetExpenseParams{
		ID:       id,
		BudgetID: budgetId,
	})
}

type BudgetWithExpenses struct {
	Budget
	Expenses []BudgetExpense
}
