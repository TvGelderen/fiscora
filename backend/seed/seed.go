package seed

import (
	"context"
	"database/sql"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/tvgelderen/budget-buddy/database"
	"github.com/tvgelderen/budget-buddy/types"
)

var userId uuid.UUID

func Seed(conn *sql.DB) {
	db := database.New(conn)

	log.Info("Seeding database")

    userId, _ = uuid.NewUUID()
    err := createDemoUser(db)
	if err != nil {
		log.Fatal("Error getting user from db: ", err.Error())
	}

	err = createTransactions(db)
	if err != nil {
		log.Fatal("Error creating test transactions: ", err.Error())
	}
}

func createDemoUser(db *database.Queries) error {
	_, err := db.CreateUser(context.Background(), database.CreateUserParams{
		ID:         userId,
		Provider:   "demo",
		ProviderID: "demo",
		Username:   "demo",
		Email:      "demo",
		Created:    time.Now().UTC(),
		Updated:    time.Now().UTC(),
	})

	return err
}

func createTransactions(db *database.Queries) error {
	for _, transaction := range transactions {
        time := randomTime()

		transaction.UserID = userId
        transaction.Created = time
        transaction.Updated = time

        if !transaction.Recurring {
            transaction.StartDate = randomTime()
            transaction.EndDate = transaction.StartDate
        }

		_, err := db.CreateTransaction(context.Background(), transaction)
		if err != nil {
			return err
		}
	}

	return nil
}

func randomTime() time.Time {
    return time.Now().UTC().AddDate(0, 0, -rand.Intn(60))
}

var transactions = []database.CreateTransactionParams{
	{
		UserID:       userId,
		Amount:       "4321",
		Description:  "Salary",
		Incoming:     true,
		Type:         types.IncomeTypeSalary,
		Recurring:    true,
		StartDate:    time.Date(2024, 7, 25, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2025, 6, 25, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "128",
		Description:  "Weekly income",
		Incoming:     true,
		Type:         types.IncomeTypePassive,
		Recurring:    true,
		StartDate:    time.Date(2024, 5, 12, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2025, 4, 12, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalWeekly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "632",
		Description:  "Mortgage",
		Incoming:     false,
		Type:         types.ExpenseTypeMortgage,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 1, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "215",
		Description:  "Utilities",
		Incoming:     false,
		Type:         types.ExpenseTypeUtilities,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 1, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "64",
		Description:  "Internet",
		Incoming:     false,
		Type:         types.ExpenseTypeUtilities,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 6, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 6, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "156",
		Description:  "HOA",
		Incoming:     false,
		Type:         types.ExpenseTypeUtilities,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 1, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "124",
		Description:  "Health insurance",
		Incoming:     false,
		Type:         types.ExpenseTypeInsurance,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 10, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "54",
		Description:  "Various subscriptions",
		Incoming:     false,
		Type:         types.ExpenseTypeSubscriptions,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 10, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "123",
		Description:  "Groceries",
		Incoming:     false,
		Type:         types.ExpenseTypeGroceries,
		Recurring:    false,
		Interval:     sql.NullString{Valid: false},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "84",
		Description:  "Groceries",
		Incoming:     false,
		Type:         types.ExpenseTypeGroceries,
		Recurring:    false,
		Interval:     sql.NullString{Valid: false},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "72",
		Description:  "Groceries",
		Incoming:     false,
		Type:         types.ExpenseTypeGroceries,
		Recurring:    false,
		Interval:     sql.NullString{Valid: false},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "95",
		Description:  "Groceries",
		Incoming:     false,
		Type:         types.ExpenseTypeGroceries,
		Recurring:    false,
		Interval:     sql.NullString{Valid: false},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "38",
		Description:  "Groceries",
		Incoming:     false,
		Type:         types.ExpenseTypeGroceries,
		Recurring:    false,
		Interval:     sql.NullString{Valid: false},
		DaysInterval: sql.NullInt32{Valid: false},
	},
}
