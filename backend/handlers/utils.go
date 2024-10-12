package handlers

import (
	"context"
	"database/sql"
	"fmt"
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

func getTransactionsFromDB(ctx context.Context, incomeParam string, userId uuid.UUID, dateRange types.DateRange, db *database.Queries) ([]database.Transaction, error) {
	income, err := strconv.ParseBool(incomeParam)

	if err != nil {
		return db.GetUserTransactionsBetweenDates(ctx, database.GetUserTransactionsBetweenDatesParams{
			UserID:    userId,
			StartDate: dateRange.End,
			EndDate:   dateRange.Start,
			Limit:     database.MaxFetchLimit,
			Offset:    0,
		})
	} else {
		if income {
			return db.GetUserIncomeTransactionsBetweenDates(ctx, database.GetUserIncomeTransactionsBetweenDatesParams{
				UserID:    userId,
				StartDate: dateRange.End,
				EndDate:   dateRange.Start,
				Limit:     database.MaxFetchLimit,
				Offset:    0,
			})
		} else {
			return db.GetUserExpenseTransactionsBetweenDates(ctx, database.GetUserExpenseTransactionsBetweenDatesParams{
				UserID:    userId,
				StartDate: dateRange.End,
				EndDate:   dateRange.Start,
				Limit:     database.MaxFetchLimit,
				Offset:    0,
			})
		}
	}
}
