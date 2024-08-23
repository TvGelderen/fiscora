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

		_, err := db.CreateTransaction(context.Background(), transaction)
		if err != nil {
			return err
		}
	}

	return nil
}

func randomTime() time.Time {
	return time.Now().UTC().AddDate(0, 0, -rand.Intn(120))
}

var transactions = []database.CreateTransactionParams{
	// Salary (Monthly, starting July 25, 2024)
	{
		UserID:       userId,
		Amount:       "3789",
		Description:  "Salary",
		Type:         types.IncomeTypeSalary,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 25, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2025, 1, 25, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	// Weekly income (Starting May 12, 2024)
	{
		UserID:       userId,
		Amount:       "128",
		Description:  "Weekly income",
		Type:         types.IncomeTypePassive,
		Recurring:    true,
		StartDate:    time.Date(2024, 5, 12, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2025, 4, 12, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalWeekly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	// Mortgage (Monthly, starting January 1, 2024)
	{
		UserID:       userId,
		Amount:       "-632",
		Description:  "Mortgage",
		Type:         types.ExpenseTypeMortgage,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 1, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	// Utilities (Monthly, starting January 1, 2024)
	{
		UserID:       userId,
		Amount:       "-215",
		Description:  "Utilities",
		Type:         types.ExpenseTypeUtilities,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 1, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	// Internet (Monthly, starting January 6, 2024)
	{
		UserID:       userId,
		Amount:       "-64",
		Description:  "Internet",
		Type:         types.ExpenseTypeUtilities,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 6, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 6, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	// HOA (Monthly, starting January 1, 2024)
	{
		UserID:       userId,
		Amount:       "-224",
		Description:  "HOA",
		Type:         types.ExpenseTypeFixed,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 1, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	// Health insurance (Monthly, starting January 10, 2024)
	{
		UserID:       userId,
		Amount:       "-124",
		Description:  "Health insurance",
		Type:         types.ExpenseTypeInsurance,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 10, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	// Various subscriptions (Monthly, starting January 10, 2024)
	{
		UserID:       userId,
		Amount:       "-54",
		Description:  "Streaming services bundle",
		Type:         types.ExpenseTypeSubscriptions,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 10, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "-12.99",
		Description:  "Music streaming service",
		Type:         types.ExpenseTypeSubscriptions,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 15, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "-9.99",
		Description:  "Cloud storage",
		Type:         types.ExpenseTypeSubscriptions,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 5, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "-29.99",
		Description:  "Online fitness membership",
		Type:         types.ExpenseTypeSubscriptions,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 1, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "-9.99",
		Description:  "Software license",
		Type:         types.ExpenseTypeSubscriptions,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 1, 15, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	{
		UserID:       userId,
		Amount:       "-4.99",
		Description:  "App subscription",
		Type:         types.ExpenseTypeSubscriptions,
		Recurring:    true,
		StartDate:    time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2032, 12, 20, 0, 0, 0, 0, time.UTC),
		Interval:     sql.NullString{String: types.TransactionIntervalMonthly, Valid: true},
		DaysInterval: sql.NullInt32{Valid: false},
	},
	// Groceries (Non-recurring, spread throughout the year)
	{
		UserID:      userId,
		Amount:      "-120",
		Description: "New Year groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-85",
		Description: "Weekly groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-100",
		Description: "Monthly stock up",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-130",
		Description: "Super Bowl party supplies",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 2, 2, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 2, 2, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-123",
		Description: "Groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 2, 18, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 2, 18, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-79",
		Description: "Groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 2, 24, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 2, 24, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-115",
		Description: "Weekly groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 3, 8, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 3, 8, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-80",
		Description: "Groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 3, 14, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 3, 14, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-95",
		Description: "Spring sale groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 3, 22, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 3, 22, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-110",
		Description: "Easter holiday groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-125",
		Description: "Groceries for BBQ party",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 4, 19, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 4, 19, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-96",
		Description: "Groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 4, 27, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 4, 27, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-130",
		Description: "Cinco de Mayo groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 5, 3, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 5, 3, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-85",
		Description: "Weekly groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 5, 17, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 5, 17, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-115",
		Description: "Memorial Day groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 5, 30, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 5, 30, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-102",
		Description: "Groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 6, 7, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 6, 7, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-90",
		Description: "Groceries for June",
		Recurring:   false,
		StartDate:   time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-95",
		Description: "Father's Day groceries",
		Recurring:   false,
		StartDate:   time.Date(2024, 6, 21, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 6, 21, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-115",
		Description: "Fourth of July groceries",
		Recurring:   false,
		StartDate:   time.Date(2024, 7, 3, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 7, 3, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-110",
		Description: "Weekly groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 7, 18, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 7, 18, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-79",
		Description: "Groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 7, 28, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 7, 28, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-150",
		Description: "Back to school groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 8, 5, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 8, 5, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-100",
		Description: "Weekly groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-78",
		Description: "Groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 8, 28, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 8, 28, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-120",
		Description: "Labor Day groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-65",
		Description: "Groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 9, 24, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 9, 24, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-130",
		Description: "Fall season groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 9, 15, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 9, 15, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-89",
		Description: "Groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 10, 02, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 10, 02, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-125",
		Description: "Groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 10, 14, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 10, 14, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-140",
		Description: "Halloween groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 10, 27, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 10, 27, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-95",
		Description: "Groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 11, 11, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 11, 11, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-160",
		Description: "Thanksgiving groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 11, 23, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 11, 23, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-86",
		Description: "Groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 12, 01, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 12, 01, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-135",
		Description: "Weekly groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 12, 10, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 12, 10, 0, 0, 0, 0, time.UTC),
	},
	{
		UserID:      userId,
		Amount:      "-165",
		Description: "Christmas holiday groceries",
		Type:        types.ExpenseTypeGroceries,
		Recurring:   false,
		StartDate:   time.Date(2024, 12, 24, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, 12, 24, 0, 0, 0, 0, time.UTC),
	},
}
