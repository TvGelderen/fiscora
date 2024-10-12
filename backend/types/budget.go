package types

import (
	"strconv"
	"time"

	"github.com/tvgelderen/fiscora/database"
)

type BaseBudget struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}

type BudgetCreateRequest struct {
	BaseBudget
}

type BudgetUpdateRequest struct {
	BaseBudget
}

type BudgetReturn struct {
	BaseBudget
	ID       string                `json:"id"`
	Created  time.Time             `json:"created"`
	Updated  time.Time             `json:"updated"`
	Expenses []BudgetExpenseReturn `json:"expenses"`
}

type BaseBudgetExpense struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	AllocatedAmount float64 `json:"allocatedAmount"`
	CurrentAmount   float64 `json:"currentAmount"`
}

type BudgetExpenseCreateRequest struct {
	BaseBudgetExpense
}

type BudgetExpenseUpdateRequest struct {
	BaseBudgetExpense
}

type BudgetExpenseReturn struct {
	BaseBudgetExpense
	ID int32 `json:"id"`
}

func ToBudget(dbModel database.Budget) BudgetReturn {
	amount, _ := strconv.ParseFloat(dbModel.Amount, 64)

	return BudgetReturn{
		ID:       dbModel.ID,
		Created:  dbModel.Created,
		Updated:  dbModel.Updated,
		Expenses: []BudgetExpenseReturn{},
		BaseBudget: BaseBudget{
			Name:        dbModel.Name,
			Description: dbModel.Description,
			Amount:      amount,
			StartDate:   dbModel.StartDate,
			EndDate:     dbModel.EndDate,
		},
	}
}

func ToBudgetWithExpenses(dbModels []database.GetBudgetsWithExpensesRow) BudgetReturn {
	amount, _ := strconv.ParseFloat(dbModels[0].Amount, 64)

	budget := BudgetReturn{
		ID:       dbModels[0].ID,
		Created:  dbModels[0].Created,
		Updated:  dbModels[0].Updated,
		Expenses: make([]BudgetExpenseReturn, len(dbModels)),
		BaseBudget: BaseBudget{
			Name:        dbModels[0].Name,
			Description: dbModels[0].Description,
			Amount:      amount,
			StartDate:   dbModels[0].StartDate,
			EndDate:     dbModels[0].EndDate,
		},
	}

	for idx, dbModel := range dbModels {
		allocatedAmount, _ := strconv.ParseFloat(dbModel.AllocatedAmount, 64)
		currentAmount, _ := strconv.ParseFloat(dbModel.CurrentAmount, 64)

		budget.Expenses[idx] = BudgetExpenseReturn{
			ID: dbModel.ID_2,
			BaseBudgetExpense: BaseBudgetExpense{
				Name:            dbModel.Name_2,
				Description:     dbModel.Description_2,
				AllocatedAmount: allocatedAmount,
				CurrentAmount:   currentAmount,
			},
		}
	}

	return budget
}
