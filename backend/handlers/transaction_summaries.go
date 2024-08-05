package handlers

import (
	"fmt"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tvgelderen/budget-buddy/database"
	"github.com/tvgelderen/budget-buddy/types"
)

func (h *APIHandler) HandleGetTransactionMonthInfo(c echo.Context) error {
	userId := getUserId(c)
    month := getMonth(c)
    year := getYear(c)

	dateRange := getMonthRange(month, year)
	transactions, err := h.DB.GetUserTransactionsBetweenDates(c.Request().Context(), database.GetUserTransactionsBetweenDatesParams{
		UserID:    userId,
		StartDate: dateRange.End,
		EndDate:   dateRange.Start,
		Limit:     database.MaxFetchLimit,
		Offset:    0,
	})
	if err != nil {
		return InternalServerError(c, fmt.Sprintf("Error getting transactions from db: %v", err.Error()))
	}

	monthInfo := types.GetMonthInfo(transactions, dateRange)

	return c.JSON(http.StatusOK, monthInfo)
}

func (h *APIHandler) HandleGetTransactionYearInfo(c echo.Context) error {
	userId := getUserId(c)
    year := getYear(c)

	yearInfo := make(map[int]types.MonthInfoReturn)

	for i := 1; i < 13; i++ {
		dateRange := getMonthRange(i, year)
		transactions, err := h.DB.GetUserTransactionsBetweenDates(c.Request().Context(), database.GetUserTransactionsBetweenDatesParams{
			UserID:    userId,
			StartDate: dateRange.End,
			EndDate:   dateRange.Start,
			Limit:     database.MaxFetchLimit,
			Offset:    0,
		})
		if err != nil {
			return InternalServerError(c, fmt.Sprintf("Error getting transactions from db: %v", err.Error()))
		}

		monthInfo := types.GetMonthInfo(transactions, dateRange)

		yearInfo[i] = monthInfo
	}

	return c.JSON(http.StatusOK, yearInfo)
}

func (h *APIHandler) HandleGetExpensePerType(c echo.Context) error {
	userId := getUserId(c)
    month := getMonth(c)
    year := getYear(c)
	dateRange := getMonthRange(month, year)

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
