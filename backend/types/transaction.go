package types

import (
	"database/sql"
	"math"
	"strconv"
	"time"

	"github.com/tvgelderen/fiscora/repository"
)

type BaseTransaction struct {
	Description  string     `json:"description"`
	Amount       float64    `json:"amount"`
	Type         string     `json:"type"`
	StartDate    NullTime   `json:"startDate"`
	EndDate      NullTime   `json:"endDate"`
	Interval     NullString `json:"interval"`
	DaysInterval NullInt    `json:"daysInterval"`
}

type TransactionCreateRequest struct {
	BaseTransaction
}

type TransactionUpdateRequest struct {
	BaseTransaction
}

type TransactionReturn struct {
	BaseTransaction
	ID                     int32     `json:"id"`
	RecurringTransactionID NullInt   `json:"recurringTransactionId"`
	Date                   time.Time `json:"date"`
	Created                time.Time `json:"created"`
	Updated                time.Time `json:"updated"`
}

type MonthInfoReturn struct {
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
}

type DateRange struct {
	Start time.Time
	End   time.Time
}

func ToReturnTransaction(transaction repository.FullTransaction) TransactionReturn {
	amount, _ := strconv.ParseFloat(transaction.Amount, 64)

	result := TransactionReturn{
		ID:                     transaction.ID,
		RecurringTransactionID: NewNullInt(transaction.RecurringTransactionID),
		Date:                   transaction.Date,
		BaseTransaction: BaseTransaction{
			Amount:      amount,
			Description: transaction.Description,
			Type:        transaction.Type,
		},
	}

	if transaction.RecurringTransactionID.Valid {
		result.Created = transaction.RecurringCreated.Time
		result.Updated = transaction.RecurringUpdated.Time
		result.BaseTransaction.StartDate = NewNullTime(transaction.StartDate)
		result.BaseTransaction.EndDate = NewNullTime(transaction.EndDate)
		result.BaseTransaction.Interval = NewNullString(transaction.Interval)
		result.BaseTransaction.DaysInterval = NewNullInt(transaction.DaysInterval)
	} else {
		result.Created = transaction.Created
		result.Updated = transaction.Created
		result.BaseTransaction.StartDate = NewNullTime(sql.NullTime{Valid: false})
		result.BaseTransaction.EndDate = NewNullTime(sql.NullTime{Valid: false})
		result.BaseTransaction.Interval = NewNullString(sql.NullString{Valid: false})
		result.BaseTransaction.DaysInterval = NewNullInt(sql.NullInt32{Valid: false})
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
