package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tvgelderen/budget-buddy/database"
	"github.com/tvgelderen/budget-buddy/types"
)

func (h *APIHandler) HandleGetExpensePerType(c echo.Context) error {
	userId := GetUserId(c)

	monthParam := c.QueryParam("month")
	yearParam := c.QueryParam("year")

	month, err := strconv.ParseInt(monthParam, 10, 16)
	if err != nil {
		month = int64(time.Now().Month())
	}
	year, err := strconv.ParseInt(yearParam, 10, 16)
	if err != nil {
		year = int64(time.Now().Year())
	}

	dateRange := getMonthRange(int(month), int(year))

	dbTransactions, err := h.DB.GetUserExpenseTransactionsBetweenDates(c.Request().Context(), database.GetUserExpenseTransactionsBetweenDatesParams{
		UserID:    userId,
		StartDate: dateRange.End,
		EndDate:   dateRange.Start,
		Limit:     database.MaxFetchLimit,
		Offset:    0,
	})
	if err != nil {
		if database.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		return InternalServerError(c, fmt.Sprintf("Error getting transactions from db: %v", err.Error()))
	}

	transactions := types.ToTransactions(dbTransactions, dateRange)

	expenses := make(map[string]float64, len(transactions))
	for _, expenseType := range types.ExpenseTypes {
		expenses[expenseType] = 0
	}

	for _, transaction := range transactions {
		expenses[transaction.Type] += math.Abs(transaction.Amount)
	}

	for key := range expenses {
		if expenses[key] == 0 {
			delete(expenses, key)
		}
	}

	return c.JSON(http.StatusOK, expenses)
}
