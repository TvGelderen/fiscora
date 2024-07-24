package types

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tvgelderen/budget-buddy/database"
)

type BaseTransaction struct {
	Amount       float64    `json:"amount"`
	Description  string     `json:"description"`
	Incoming     bool       `json:"incoming"`
	Type         string     `json:"type"`
	StartDate    time.Time  `json:"startDate"`
	EndDate      time.Time  `json:"endDate"`
	Recurring    bool       `json:"recurring"`
	Interval     NullString `json:"interval"`
	DaysInterval NullInt    `json:"daysInterval"`
}

type TransactionCreateRequest struct {
	BaseTransaction
}

type TransactionReturn struct {
	ID   int64     `json:"id"`
	Date time.Time `json:"date"`
	BaseTransaction
}

type DateRange struct {
	Start time.Time
	End   time.Time
}

func ToTransaction(dbModel database.Transaction, date time.Time) TransactionReturn {
	amount, _ := strconv.ParseFloat(dbModel.Amount, 64)
	return TransactionReturn{
		ID:   dbModel.ID,
		Date: date,
		BaseTransaction: BaseTransaction{
			Amount:       amount,
			Description:  dbModel.Description,
			Incoming:     dbModel.Incoming,
			Type:         dbModel.Type,
			StartDate:    dbModel.StartDate,
			Recurring:    dbModel.Recurring,
			EndDate:      dbModel.EndDate,
			Interval:     NewNullString(dbModel.Interval),
			DaysInterval: NewNullInt(dbModel.DaysInterval),
		},
	}
}

func AddDate(transaction TransactionReturn, date time.Time) TransactionReturn {
	return TransactionReturn{
		ID:   transaction.ID,
		Date: date,
		BaseTransaction: BaseTransaction{
			Amount:       transaction.Amount,
			Description:  transaction.Description,
			Incoming:     transaction.Incoming,
			Type:         transaction.Type,
			StartDate:    transaction.StartDate,
			Recurring:    transaction.Recurring,
			EndDate:      transaction.EndDate,
			Interval:     transaction.Interval,
			DaysInterval: transaction.DaysInterval,
		},
	}
}

func ToTransactions(dbModels []database.Transaction, dateRange DateRange) []TransactionReturn {
	transactions := []TransactionReturn{}

	for _, transaction := range dbModels {
		if !transaction.Recurring {
            transactions = append(transactions, ToTransaction(transaction, transaction.StartDate))
			continue
		}

		date := transaction.StartDate

		switch transaction.Interval.String {
		case TransactionIntervalDaily:
			if date.Before(dateRange.Start) {
                date = addDays(date, daysBetween(date, dateRange.Start))
			}
			for date.Before(dateRange.End) && date.Before(transaction.EndDate) {
				transactions = append(transactions, ToTransaction(transaction, date))
                date = addDays(date, 1)
			}
		case TransactionIntervalWeekly:
			if date.Before(dateRange.Start) {
                date = addDays(date, int(daysBetween(date, dateRange.Start) / 7))
			}
			for date.Before(dateRange.End) && date.Before(transaction.EndDate) {
				transactions = append(transactions, ToTransaction(transaction, date))
                date = addWeek(date)
			}
		case TransactionIntervalMonthly:
			if date.Before(dateRange.Start) {
                date = addMonths(date, monthsBetween(date, dateRange.Start))
			}
			for date.Before(dateRange.End) && date.Before(transaction.EndDate) {
				transactions = append(transactions, ToTransaction(transaction, date))
                date = addMonths(date, 1)
			}
        case TransactionIntervalOther:
            if transaction.DaysInterval.Int32 == 0 {
                continue
            }
            if date.Before(dateRange.Start) {
                date = addDays(date, int(transaction.DaysInterval.Int32))
            }
			for date.Before(dateRange.End) && date.Before(transaction.EndDate) {
				transactions = append(transactions, ToTransaction(transaction, date))
                date = addDays(date, int(transaction.DaysInterval.Int32))
			}
		}
	}

	return transactions
}

func addDays(date time.Time, days int) time.Time {
    return date.AddDate(0, 0, days)
}

func addWeek(date time.Time) time.Time {
    return date.AddDate(0, 0, 7)
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
    if end.Day() < start.Day() {
        totalMonths--
    }
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
)

var IncomeTypes = []string{
	IncomeTypeSalary,
	IncomeTypePassive,
	IncomeTypeCapitalGains,
	IncomeTypeDividend,
	IncomeTypeGovernmentPayment,
}

// Expens type
const (
	ExpenseTypeMortgage  string = "Mortgage"
	ExpenseTypeRent             = "Rent"
	ExpenseTypeUtilities        = "Utilities"
	ExpenseTypeFixed            = "Fixed"
	ExpenseTypeGroceries        = "Groceries"
	ExpenseTypeInsurance        = "Insurance"
	ExpenseTypeTravel           = "Travel"
	ExpenseTypeTaxes            = "Taxes"
	ExpenseTypeInterest         = "Interest"
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
}
