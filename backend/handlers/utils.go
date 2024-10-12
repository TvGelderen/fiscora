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

func getTransactionsFromDB(ctx context.Context, incomeParam string, userId uuid.UUID, dateRange types.DateRange, db *database.Queries) ([]database.Transaction, error) {
	income, err := strconv.ParseBool(incomeParam)

	if err != nil {
		return db.GetTransactionsBetweenDates(ctx, database.GetTransactionsBetweenDatesParams{
			UserID:    userId,
			StartDate: dateRange.End,
			EndDate:   dateRange.Start,
			Limit:     database.MaxFetchLimit,
			Offset:    0,
		})
	} else {
		if income {
			return db.GetIncomingTransactionsBetweenDates(ctx, database.GetIncomingTransactionsBetweenDatesParams{
				UserID:    userId,
				StartDate: dateRange.End,
				EndDate:   dateRange.Start,
				Limit:     database.MaxFetchLimit,
				Offset:    0,
			})
		} else {
			return db.GetOutgoingTransactionsBetweenDates(ctx, database.GetOutgoingTransactionsBetweenDatesParams{
				UserID:    userId,
				StartDate: dateRange.End,
				EndDate:   dateRange.Start,
				Limit:     database.MaxFetchLimit,
				Offset:    0,
			})
		}
	}
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
	dbBudgetExpenses, err := db.GetBudgetsWithExpenses(ctx, database.GetBudgetsWithExpensesParams{
		UserID: userId,
		Limit:  database.MaxFetchLimit,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}

	dbBudgetMap := make(map[string][]database.GetBudgetsWithExpensesRow, len(dbBudgets))

	for _, budget := range dbBudgets {
		dbBudgetMap[budget.ID] = []database.GetBudgetsWithExpensesRow{}
	}

	for _, budgetExpense := range dbBudgetExpenses {
		dbBudgetMap[budgetExpense.ID] = append(dbBudgetMap[budgetExpense.ID], budgetExpense)
	}

	budgetCount := 0
	budgets := make([]types.BudgetReturn, len(dbBudgets))

	for _, dbBudget := range dbBudgets {
		if len(dbBudgetMap[dbBudget.ID]) > 0 {
			budgets[budgetCount] = types.ToBudgetWithExpenses(dbBudgetMap[dbBudget.ID])
		} else {
			budgets[budgetCount] = types.ToBudget(dbBudget)
		}
		budgetCount++
	}

	return budgets, nil
}
