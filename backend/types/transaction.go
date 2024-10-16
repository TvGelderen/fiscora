package types

import (
	"database/sql"
	"math"
	"strconv"
	"time"

	"github.com/google/uuid"
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

func ToTransaction(transaction repository.FullTransaction) TransactionReturn {
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

func CreateRecurringTransactions(recurringTransactionId int32, transaction BaseTransaction, userId uuid.UUID) []repository.CreateTransactionParams {
	if !transaction.StartDate.Valid ||
		!transaction.EndDate.Valid ||
		!transaction.Interval.Valid {
		return []repository.CreateTransactionParams{}
	}

	date := transaction.StartDate.Time
	endDate := transaction.EndDate.Time
	amount := strconv.FormatFloat(transaction.Amount, 'f', -1, 64)

	createParams := []repository.CreateTransactionParams{}

	for date.Before(endDate) {
		createParams = append(createParams, repository.CreateTransactionParams{
			UserID:                 userId,
			RecurringTransactionID: sql.NullInt32{Valid: true, Int32: recurringTransactionId},
			Amount:                 amount,
			Description:            transaction.Description,
			Type:                   transaction.Type,
			Date:                   date,
		})

		switch transaction.Interval.String {
		case TransactionIntervalDaily:
			date = addDays(date, 1)
		case TransactionIntervalWeekly:
			date = addWeeks(date, 1)
		case TransactionIntervalMonthly:
			date = addMonths(date, 1)
		case TransactionIntervalOther:
			date = addDays(date, int(transaction.DaysInterval.Int32))
		}
	}

	return createParams
}

func GetMonthInfo(transactions []repository.Transaction) MonthInfoReturn {
	var income float64 = 0
	var expense float64 = 0

	for _, transaction := range transactions {
		amount := getAmount(transaction)
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

func getAmount(transaction repository.Transaction) float64 {
	amount, _ := strconv.ParseFloat(transaction.Amount, 64)
	return amount
}

func addDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

func addWeeks(date time.Time, weeks int) time.Time {
	return date.AddDate(0, 0, weeks*7)
}

func addMonths(date time.Time, months int) time.Time {
	return date.AddDate(0, months, 0)
}

func daysBetween(start, end time.Time) int {
	if start.After(end) {
		start, end = end, start
	}
	duration := end.Sub(start)
	return int(duration.Hours() / 24)
}

func monthsBetween(start, end time.Time) int {
	if start.After(end) {
		start, end = end, start
	}
	years := end.Year() - start.Year()
	months := int(end.Month()) - int(start.Month())
	totalMonths := (years * 12) + months
	return totalMonths
}

// Transaction interval
const (
	TransactionIntervalDaily   string = "Daily"
	TransactionIntervalWeekly         = "Weekly"
	TransactionIntervalMonthly        = "Monthly"
	TransactionIntervalOther          = "Other"
)

var TransactionIntervals = []string{
	TransactionIntervalDaily,
	TransactionIntervalWeekly,
	TransactionIntervalMonthly,
	TransactionIntervalOther,
}

// Income type
const (
	IncomeTypeSalary            string = "Salary"
	IncomeTypePassive                  = "Passive"
	IncomeTypeCapitalGains             = "Capital Gains"
	IncomeTypeDividend                 = "Dividend"
	IncomeTypeGovernmentPayment        = "Government Payment"
	IncomeTypeOther                    = "Other"
)

var IncomeTypes = []string{
	IncomeTypeSalary,
	IncomeTypePassive,
	IncomeTypeCapitalGains,
	IncomeTypeDividend,
	IncomeTypeGovernmentPayment,
	IncomeTypeOther,
}

// Expense type
const (
	ExpenseTypeMortgage      string = "Mortgage"
	ExpenseTypeRent                 = "Rent"
	ExpenseTypeUtilities            = "Utilities"
	ExpenseTypeFixed                = "Fixed"
	ExpenseTypeGroceries            = "Groceries"
	ExpenseTypeInsurance            = "Insurance"
	ExpenseTypeTravel               = "Travel"
	ExpenseTypeTaxes                = "Taxes"
	ExpenseTypeInterest             = "Interest"
	ExpenseTypeSubscriptions        = "Subscriptions"
	ExpenseTypeOther                = "Other"
)

var ExpenseTypes = []string{
	ExpenseTypeMortgage,
	ExpenseTypeRent,
	ExpenseTypeUtilities,
	ExpenseTypeFixed,
	ExpenseTypeGroceries,
	ExpenseTypeInsurance,
	ExpenseTypeTravel,
	ExpenseTypeTaxes,
	ExpenseTypeInterest,
	ExpenseTypeSubscriptions,
	ExpenseTypeOther,
}
