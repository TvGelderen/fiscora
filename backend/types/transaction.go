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
	EndDate      time.Time  `json:"endDate"`
	Recurring    bool       `json:"recurring"`
	Interval     NullString `json:"interval"`
	DaysInterval NullInt    `json:"daysInterval"`
}

type TransactionCreateRequest struct {
	BaseTransaction
}

type TransactionReturn struct {
	ID int64 `json:"id"`
	BaseTransaction
}

type DateRange struct {
	Start time.Time
	End   time.Time
}

func ToTransaction(dbModel database.Transaction) TransactionReturn {
	amount, _ := strconv.ParseFloat(dbModel.Amount, 64)
	return TransactionReturn{
		ID: dbModel.ID,
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

func ToTransactions(dbModels []database.Transaction) []TransactionReturn {
	transations := make([]TransactionReturn, len(dbModels))

	for idx, model := range dbModels {
		transations[idx] = ToTransaction(model)
	}

	return transations
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
