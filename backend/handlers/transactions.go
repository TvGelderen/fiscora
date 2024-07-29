package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tvgelderen/budget-buddy/database"
	"github.com/tvgelderen/budget-buddy/types"
)

func (h *APIHandler) HandleGetTransactionIntervals(c echo.Context) error {
	intervals := make([]string, len(types.TransactionIntervals))

	for idx, interval := range types.TransactionIntervals {
		intervals[idx] = interval
	}

	return c.JSON(http.StatusOK, intervals)
}

func (h *APIHandler) HandleGetIncomeTypes(c echo.Context) error {
	incomeTypes := make([]string, len(types.IncomeTypes))

	for idx, incomeType := range types.IncomeTypes {
		incomeTypes[idx] = incomeType
	}

	return c.JSON(http.StatusOK, incomeTypes)
}

func (h *APIHandler) HandleGetExpenseTypes(c echo.Context) error {
	expenseTypes := make([]string, len(types.ExpenseTypes))

	for idx, expenseType := range types.ExpenseTypes {
		expenseTypes[idx] = expenseType
	}

	return c.JSON(http.StatusOK, expenseTypes)
}

func (h *APIHandler) HandleGetTransactions(c echo.Context) error {
	userId := GetUserId(c)

	skipParam := c.QueryParam("skip")
	takeParam := c.QueryParam("take")
	monthParam := c.QueryParam("month")
	yearParam := c.QueryParam("year")
	incomeParam := c.QueryParam("income")

	skip, err := strconv.ParseInt(skipParam, 10, 32)
	if err != nil {
		skip = 0
	}
	take, err := strconv.ParseInt(takeParam, 10, 32)
	if err != nil {
		take = database.DefaultFetchLimit
	}
	month, err := strconv.ParseInt(monthParam, 10, 16)
	if err != nil {
		month = int64(time.Now().Month())
	}
	year, err := strconv.ParseInt(yearParam, 10, 16)
	if err != nil {
		year = int64(time.Now().Year())
	}
	income, err := strconv.ParseBool(incomeParam)

	dateRange := getMonthRange(int(month), int(year))

	var transactions []database.Transaction
	if err != nil {
		transactions, err = h.DB.GetUserTransactionsBetweenDates(c.Request().Context(), database.GetUserTransactionsBetweenDatesParams{
			UserID:    userId,
			StartDate: dateRange.End,
			EndDate:   dateRange.Start,
			Limit:     int32(take),
			Offset:    int32(skip),
		})
	} else {
		if income {
			transactions, err = h.DB.GetUserIncomeTransactionsBetweenDates(c.Request().Context(), database.GetUserIncomeTransactionsBetweenDatesParams{
				UserID:    userId,
				StartDate: dateRange.End,
				EndDate:   dateRange.Start,
				Limit:     int32(take),
				Offset:    int32(skip),
			})
		} else {
			transactions, err = h.DB.GetUserExpenseTransactionsBetweenDates(c.Request().Context(), database.GetUserExpenseTransactionsBetweenDatesParams{
				UserID:    userId,
				StartDate: dateRange.End,
				EndDate:   dateRange.Start,
				Limit:     int32(take),
				Offset:    int32(skip),
			})
		}
	}
	if err != nil {
		if database.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		return InternalServerError(c, fmt.Sprintf("Error getting transactions from db: %v", err.Error()))
	}

	return c.JSON(http.StatusOK, types.ToTransactions(transactions, dateRange))
}

func (h *APIHandler) HandleCreateTransaction(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	transaction := types.TransactionCreateRequest{}
	err := decoder.Decode(&transaction)
	if err != nil {
		log.Errorf("Error decoding request body: %v", err.Error())
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	// TODO: validate transaction object

	userId := GetUserId(c)

	_, err = h.DB.CreateTransaction(c.Request().Context(), database.CreateTransactionParams{
		UserID:       userId,
		Amount:       strconv.FormatFloat(transaction.Amount, 'f', -1, 64),
		Description:  transaction.Description,
		Incoming:     transaction.Incoming,
		Type:         transaction.Type,
		Recurring:    transaction.Recurring,
		StartDate:    transaction.StartDate,
		EndDate:      transaction.EndDate,
		Interval:     transaction.Interval.NullString,
		DaysInterval: transaction.DaysInterval.NullInt32,
		Created:      time.Now().UTC(),
		Updated:      time.Now().UTC(),
	})
	if err != nil {
		return InternalServerError(c, fmt.Sprintf("Error creating transaction: %v", err.Error()))
	}

	return c.String(http.StatusCreated, "Transaction created successfully")
}

func (h *APIHandler) HandleDeleteTransaction(c echo.Context) error {
	userId := GetUserId(c)
    transactionIdParam := c.Param("id")

    transactionId, err := strconv.ParseInt(transactionIdParam, 10, 64)
    if err != nil {
        log.Errorf("Error parsing transaction id from request: %v", err.Error())
        return c.NoContent(http.StatusBadRequest)
    }

    err = h.DB.DeleteTransaction(c.Request().Context(), database.DeleteTransactionParams{
        UserID: userId,
        ID: transactionId,
    })
    if err != nil {
        return InternalServerError(c, fmt.Sprintf("Error deleting transaction: %v", err.Error()))
    }

    return c.NoContent(204)
}

func (h *APIHandler) HandleGetTransactionMonthInfo(c echo.Context) error {
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
	transactions, err := h.DB.GetUserTransactionsBetweenDates(c.Request().Context(), database.GetUserTransactionsBetweenDatesParams{
		UserID:    userId,
		StartDate: dateRange.End,
		EndDate:   dateRange.Start,
		Limit:     database.MaxFetchLimit,
		Offset:    0,
	})

	amounts := types.ToTransactionAmounts(transactions, dateRange)

	var income float64
	var expense float64

	for _, amount := range amounts {
		if amount > 0 {
			income += amount
		} else {
			expense += amount
		}
	}

	expense *= -1

	return c.JSON(http.StatusOK, types.MonthInfoReturn{
		Income:  income,
		Expense: expense,
	})
}

func getMonthRange(month int, year int) types.DateRange {
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, -1)
	return types.DateRange{
		Start: start,
		End:   end,
	}
}
