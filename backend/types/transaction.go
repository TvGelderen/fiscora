package types

import (
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
	Recurring    bool       `json:"recurring"`
	EndDate      NullTime   `json:"endDate"`
	Interval     NullString `json:"interval"`
	DaysInterval NullInt    `json:"daysInterval"`
}

type TransactionCreateRequest struct {
	BaseTransaction
}

type TransactionReturn struct {
	ID int32 `json:"id"`
	BaseTransaction
}

func ToTransaction(transaction database.Transaction) TransactionReturn {
	amount, _ := strconv.ParseFloat(transaction.Amount, 64)
	return TransactionReturn{
		ID: transaction.ID,
		BaseTransaction: BaseTransaction{
			Amount:       amount,
			Description:  transaction.Description,
			Incoming:     transaction.Incoming,
			Type:         transaction.Type,
			StartDate:    transaction.StartDate,
			Recurring:    transaction.Recurring,
			EndDate:      NewNullTime(transaction.EndDate),
			Interval:     NewNullString(transaction.Interval),
			DaysInterval: NewNullInt(transaction.DaysInterval),
		},
	}
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
