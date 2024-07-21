package models

type Transaction struct {
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
