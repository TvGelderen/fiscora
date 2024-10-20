package types

import (
	"math"
	"strconv"
	"time"

	"github.com/tvgelderen/fiscora/repository"
)

type TransactionForm struct {
	Description  string     `json:"description"`
	Amount       float64    `json:"amount"`
	Type         string     `json:"type"`
	Recurring    bool       `json:"recurring"`
	StartDate    NullTime   `json:"startDate"`
	EndDate      NullTime   `json:"endDate"`
	Interval     NullString `json:"interval"`
	DaysInterval NullInt    `json:"daysInterval"`
}

type TransactionReturn struct {
	ID          int32                 `json:"id"`
	Description string                `json:"description"`
	Amount      float64               `json:"amount"`
	Type        string                `json:"type"`
	Date        time.Time             `json:"date"`
	Created     time.Time             `json:"created"`
	Updated     time.Time             `json:"updated"`
	Recurring   *TransactionRecurring `json:"recurring"`
	Budget      *TransactionBudget    `json:"budget"`
}

type TransactionRecurring struct {
	ID           NullInt    `json:"id"`
	StartDate    NullTime   `json:"startDate"`
	EndDate      NullTime   `json:"endDate"`
	Interval     NullString `json:"interval"`
	DaysInterval NullInt    `json:"daysInterval"`
}

type TransactionBudget struct {
	ID          NullString `json:"id"`
	Name        NullString `json:"name"`
	ExpenseName NullString `json:"expenseName"`
}

type MonthInfoReturn struct {
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
}

type DateRange struct {
	Start time.Time
	End   time.Time
}

func ToTransactionReturn(transaction repository.FullTransaction) TransactionReturn {
	amount, _ := strconv.ParseFloat(transaction.Amount, 64)

	result := TransactionReturn{
		ID:          transaction.ID,
		Description: transaction.Description,
		Amount:      amount,
		Type:        transaction.Type,
		Date:        transaction.Date,
		Recurring:   nil,
		Budget:      nil,
	}

	if transaction.RecurringTransactionID.Valid {
		result.Created = transaction.RecurringCreated.Time
		result.Updated = transaction.RecurringUpdated.Time
		result.Recurring = &TransactionRecurring{
			ID:           NewNullInt(transaction.RecurringTransactionID),
			StartDate:    NewNullTime(transaction.StartDate),
			EndDate:      NewNullTime(transaction.EndDate),
			Interval:     NewNullString(transaction.Interval),
			DaysInterval: NewNullInt(transaction.DaysInterval),
		}
	} else {
		result.Created = transaction.Created
		result.Updated = transaction.Created
	}

	if transaction.BudgetID.Valid {
		result.Budget = &TransactionBudget{
			ID:          NewNullString(transaction.BudgetID),
			Name:        NewNullString(transaction.BudgetName),
			ExpenseName: NewNullString(transaction.BudgetExpenseName),
		}
	}

	return result
}

func GetMonthInfo(amounts *[]float64) MonthInfoReturn {
	var income float64 = 0
	var expense float64 = 0

	for _, amount := range *amounts {
		if amount > 0 {
			income += amount
		} else {
			expense += amount
		}
	}

	return MonthInfoReturn{
		Income:  income,
		Expense: math.Abs(expense),
	}
}
