package types

import (
	"strconv"
	"time"

	"github.com/tvgelderen/fiscora/repository"
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

type BudgetForm struct {
	BaseBudget
	Expenses []BudgetExpenseForm `json:"expenses"`
}

type BaseBudgetExpense struct {
	Name            string  `json:"name"`
	AllocatedAmount float64 `json:"allocatedAmount"`
	CurrentAmount   float64 `json:"currentAmount"`
}

type BudgetExpenseCreateRequest struct {
	BaseBudgetExpense
}

type BudgetExpenseUpdateRequest struct {
	BaseBudgetExpense
}

type BudgetExpenseForm struct {
	BaseBudgetExpense
	ID int32 `json:"id"`
}

type BudgetExpenseReturn struct {
	BaseBudgetExpense
	ID int32 `json:"id"`
}

func ToBudget(dbModel repository.Budget) BudgetReturn {
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

func ToBudgetExpense(dbModel repository.BudgetExpense) BudgetExpenseReturn {
	allocatedAmount, _ := strconv.ParseFloat(dbModel.AllocatedAmount, 64)
	currentAmount, _ := strconv.ParseFloat(dbModel.CurrentAmount, 64)

	return BudgetExpenseReturn{
		ID: dbModel.ID,
		BaseBudgetExpense: BaseBudgetExpense{
			Name:            dbModel.Name,
			AllocatedAmount: allocatedAmount,
			CurrentAmount:   currentAmount,
		},
	}
}

const (
	BudgetTypeWeekly  string = "Weekly"
	BudgetTypeMonthly        = "Monthly"
	BudgetTypeYearly         = "Yearly"
	BudgetTypeCustom         = "Custom"
)

var BudgetTypes = []string{
	BudgetTypeWeekly,
	BudgetTypeMonthly,
	BudgetTypeYearly,
	BudgetTypeCustom,
}
