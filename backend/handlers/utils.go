package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tvgelderen/fiscora/auth"
	"github.com/tvgelderen/fiscora/database"
	"github.com/tvgelderen/fiscora/types"
)

type APIHandler struct {
	DB          *database.Queries
	AuthService *auth.AuthService
}

func NewAPIHandler(connection *sql.DB, auth *auth.AuthService) *APIHandler {
	return &APIHandler{
		DB:          database.New(connection),
		AuthService: auth,
	}
}

func InternalServerError(c echo.Context, err string) error {
	log.Error(err)
	return c.String(http.StatusInternalServerError, "Something went wrong")
}

func DataBaseQueryError(c echo.Context, err error) error {
	if database.NoRowsFound(err) {
		return c.NoContent(http.StatusNotFound)
	}
	return InternalServerError(c, fmt.Sprintf("Error getting data from db: %v", err.Error()))
}

func getUserId(c echo.Context) uuid.UUID {
	return c.Get(userIdKey).(uuid.UUID)
}

func getMonth(c echo.Context) int {
	monthParam := c.QueryParam("month")
	month, err := strconv.ParseInt(monthParam, 10, 16)
	if err != nil {
		month = int64(time.Now().Month())
	}

	return int(month)
}

func getYear(c echo.Context) int {
	yearParam := c.QueryParam("year")
	year, err := strconv.ParseInt(yearParam, 10, 16)
	if err != nil {
		year = int64(time.Now().Month())
	}

	return int(year)
}

func getMonthRange(month int, year int) types.DateRange {
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, -1)
	return types.DateRange{
		Start: start,
		End:   end,
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateRandomString(length int) string {
	str := make([]rune, length)
	for idx := range str {
		str[idx] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(str)
}

func getTransactionsFromDB(ctx context.Context, incomeParam string, userId uuid.UUID, dateRange types.DateRange, db *database.Queries) ([]types.TransactionReturn, error) {
	income, err := strconv.ParseBool(incomeParam)

	if err != nil {
		dbTransactions, err := db.GetTransactionsBetweenDates(ctx, database.GetTransactionsBetweenDatesParams{
			UserID:    userId,
			StartDate: dateRange.Start,
			EndDate:   dateRange.End,
			Limit:     database.MaxFetchLimit,
			Offset:    0,
		})
		fmt.Println(err)
		if err != nil {
			return []types.TransactionReturn{}, err
		}

		transactions := make([]types.TransactionReturn, len(dbTransactions))
		for idx, dbTransaction := range dbTransactions {
			transactions[idx] = types.ToTransaction(dbTransaction.FullTransaction)
		}

		return transactions, nil
	}

	if income {
		dbTransactions, err := db.GetIncomeTransactionsBetweenDates(ctx, database.GetIncomeTransactionsBetweenDatesParams{
			UserID:    userId,
			StartDate: dateRange.Start,
			EndDate:   dateRange.End,
			Limit:     database.MaxFetchLimit,
			Offset:    0,
		})
		if err != nil {
			return []types.TransactionReturn{}, err
		}

		transactions := make([]types.TransactionReturn, len(dbTransactions))
		for idx, dbTransaction := range dbTransactions {
			transactions[idx] = types.ToTransaction(dbTransaction.FullTransaction)
		}

		return transactions, nil
	}

	dbTransactions, err := db.GetExpenseTransactionsBetweenDates(ctx, database.GetExpenseTransactionsBetweenDatesParams{
		UserID:    userId,
		StartDate: dateRange.Start,
		EndDate:   dateRange.End,
		Limit:     database.MaxFetchLimit,
		Offset:    0,
	})
	if err != nil {
		return []types.TransactionReturn{}, err
	}

	transactions := make([]types.TransactionReturn, len(dbTransactions))
	for idx, dbTransaction := range dbTransactions {
		transactions[idx] = types.ToTransaction(dbTransaction.FullTransaction)
	}

	return transactions, nil
}

func getBudgetsFromDB(ctx context.Context, userId uuid.UUID, db *database.Queries) ([]types.BudgetReturn, error) {
	dbBudgets, err := db.GetBudgets(ctx, database.GetBudgetsParams{
		UserID: userId,
		Limit:  database.MaxFetchLimit,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}
	dbBudgetExpenses, err := db.GetBudgetsExpenses(ctx, database.GetBudgetsExpensesParams{
		UserID: userId,
		Limit:  database.MaxFetchLimit,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}

	dbBudgetMap := make(map[string][]types.BudgetExpenseReturn, len(dbBudgets))
	for _, dbBudgetExpense := range dbBudgetExpenses {
		dbBudgetMap[dbBudgetExpense.BudgetExpense.BudgetID] = append(dbBudgetMap[dbBudgetExpense.BudgetExpense.BudgetID], types.ToBudgetExpense(dbBudgetExpense.BudgetExpense))
	}

	idx := 0
	budgets := make([]types.BudgetReturn, len(dbBudgets))

	for _, dbBudget := range dbBudgets {
		budget := types.ToBudget(dbBudget)
		if val, ok := dbBudgetMap[dbBudget.ID]; ok {
			budget.Expenses = val
		} else {
			budget.Expenses = []types.BudgetExpenseReturn{}
		}
		budgets[idx] = budget
		idx++
	}

	return budgets, nil
}
