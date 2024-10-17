package repository

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type ITransactionRepository interface {
	GetBetweenDates(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) (*[]FullTransaction, error)
	GetIncomeBetweenDates(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) (*[]FullTransaction, error)
	GetExpenseBetweenDates(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) (*[]FullTransaction, error)

	GetAmountsBetweenDates(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) (*[]float64, error)
	GetIncomeAmountsBetweenDates(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) (*[]TypeAmount, error)
	GetExpenseAmountsBetweenDates(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) (*[]TypeAmount, error)

	Add(ctx context.Context, params CreateTransactionParams) (*Transaction, error)
	Update(ctx context.Context, params UpdateTransactionParams) error
	Remove(ctx context.Context, id int32, userId uuid.UUID) error

	AddRecurring(ctx context.Context, params CreateRecurringTransactionParams, amount float64, description string, transactionType string) error
	UpdateRecurring(ctx context.Context, params UpdateRecurringTransactionParams) error
	RemoveRecurring(ctx context.Context, id int32, userId uuid.UUID) error
}

type TransactionRepository struct {
	db *sql.DB
}

func CreateTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (repository *TransactionRepository) GetBetweenDates(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) (*[]FullTransaction, error) {
	db := New(repository.db)
	transactions, err := db.GetTransactionsBetweenDates(ctx, GetTransactionsBetweenDatesParams{
		UserID:    userId,
		StartDate: start,
		EndDate:   end,
		Limit:     MaxFetchLimit,
		Offset:    0,
	})
	if err != nil {
		return nil, err
	}

	fullTransactions := make([]FullTransaction, len(transactions))
	for idx, transaction := range transactions {
		fullTransactions[idx] = transaction.FullTransaction
	}

	return &fullTransactions, err
}

func (repository *TransactionRepository) GetIncomeBetweenDates(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) (*[]FullTransaction, error) {
	db := New(repository.db)
	transactions, err := db.GetIncomeTransactionsBetweenDates(ctx, GetIncomeTransactionsBetweenDatesParams{
		UserID:    userId,
		StartDate: start,
		EndDate:   end,
		Limit:     MaxFetchLimit,
		Offset:    0,
	})
	if err != nil {
		return nil, err
	}

	fullTransactions := make([]FullTransaction, len(transactions))
	for idx, transaction := range transactions {
		fullTransactions[idx] = transaction.FullTransaction
	}

	return &fullTransactions, err
}

func (repository *TransactionRepository) GetExpenseBetweenDates(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) (*[]FullTransaction, error) {
	db := New(repository.db)
	transactions, err := db.GetExpenseTransactionsBetweenDates(ctx, GetExpenseTransactionsBetweenDatesParams{
		UserID:    userId,
		StartDate: start,
		EndDate:   end,
		Limit:     MaxFetchLimit,
		Offset:    0,
	})
	if err != nil {
		return nil, err
	}

	fullTransactions := make([]FullTransaction, len(transactions))
	for idx, transaction := range transactions {
		fullTransactions[idx] = transaction.FullTransaction
	}

	return &fullTransactions, err
}

func (repository *TransactionRepository) GetAmountsBetweenDates(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) (*[]float64, error) {
	db := New(repository.db)
	amounts, err := db.GetTransactionAmountsBetweenDates(ctx, GetTransactionAmountsBetweenDatesParams{
		UserID:    userId,
		StartDate: start,
		EndDate:   end,
	})
	if err != nil {
		return nil, err
	}

	floats := make([]float64, len(amounts))
	for idx, typeAmount := range amounts {
		value, err := strconv.ParseFloat(typeAmount, 64)
		if err != nil {
			value = 0
		}

		floats[idx] = value
	}

	return &floats, nil
}

func (repository *TransactionRepository) GetIncomeAmountsBetweenDates(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) (*[]TypeAmount, error) {
	db := New(repository.db)
	typeAmounts, err := db.GetIncomeTransactionAmountsBetweenDates(ctx, GetIncomeTransactionAmountsBetweenDatesParams{
		UserID:    userId,
		StartDate: start,
		EndDate:   end,
	})
	if err != nil {
		return nil, err
	}

	returnValues := make([]TypeAmount, len(typeAmounts))
	for idx, typeAmount := range typeAmounts {
		value, err := strconv.ParseFloat(typeAmount.Amount, 64)
		if err != nil {
			value = 0
		}

		returnValues[idx] = TypeAmount{
			Type:   typeAmount.Type,
			Amount: value,
		}
	}

	return &returnValues, nil
}

func (repository *TransactionRepository) GetExpenseAmountsBetweenDates(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) (*[]TypeAmount, error) {
	db := New(repository.db)
	typeAmounts, err := db.GetExpenseTransactionAmountsBetweenDates(ctx, GetExpenseTransactionAmountsBetweenDatesParams{
		UserID:    userId,
		StartDate: start,
		EndDate:   end,
	})
	if err != nil {
		return nil, err
	}

	returnValues := make([]TypeAmount, len(typeAmounts))
	for idx, typeAmount := range typeAmounts {
		value, err := strconv.ParseFloat(typeAmount.Amount, 64)
		if err != nil {
			value = 0
		}

		returnValues[idx] = TypeAmount{
			Type:   typeAmount.Type,
			Amount: value,
		}
	}

	return &returnValues, nil
}

func (repository *TransactionRepository) Add(ctx context.Context, params CreateTransactionParams) (*Transaction, error) {
	db := New(repository.db)
	transaction, err := db.CreateTransaction(ctx, params)
	return &transaction, err
}

func (repository *TransactionRepository) Update(ctx context.Context, params UpdateTransactionParams) error {
	db := New(repository.db)
	return db.UpdateTransaction(ctx, params)
}

func (repository *TransactionRepository) Remove(ctx context.Context, id int32, userId uuid.UUID) error {
	db := New(repository.db)
	return db.DeleteTransaction(ctx, DeleteTransactionParams{
		ID:     id,
		UserID: userId,
	})
}

func (repository *TransactionRepository) AddRecurring(ctx context.Context, params CreateRecurringTransactionParams, amount float64, description string, transactionType string) error {
	db := New(repository.db)
	recurringTransaction, err := db.CreateRecurringTransaction(ctx, params)
	if err != nil {
		return err
	}

	date := recurringTransaction.StartDate
	endDate := recurringTransaction.EndDate
	amountString := strconv.FormatFloat(amount, 'f', -1, 64)

	for date.Before(endDate) {
		repository.Add(ctx, CreateTransactionParams{
			UserID:                 recurringTransaction.UserID,
			RecurringTransactionID: sql.NullInt32{Valid: true, Int32: recurringTransaction.ID},
			Amount:                 amountString,
			Description:            description,
			Type:                   transactionType,
			Date:                   date,
		})

		switch recurringTransaction.Interval {
		case TransactionIntervalDaily:
			date = addDays(date, 1)
		case TransactionIntervalWeekly:
			date = addWeeks(date, 1)
		case TransactionIntervalMonthly:
			date = addMonths(date, 1)
		case TransactionIntervalOther:
			date = addDays(date, int(recurringTransaction.DaysInterval.Int32))
		}
	}

	return err
}

func (repository *TransactionRepository) UpdateRecurring(ctx context.Context, params UpdateRecurringTransactionParams) error {
	db := New(repository.db)
	return db.UpdateRecurringTransaction(ctx, params)
}

func (repository *TransactionRepository) RemoveRecurring(ctx context.Context, id int32, userId uuid.UUID) error {
	db := New(repository.db)
	return db.DeleteRecurringTransaction(ctx, DeleteRecurringTransactionParams{
		ID:     id,
		UserID: userId,
	})
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

type TypeAmount struct {
	Type   string
	Amount float64
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
