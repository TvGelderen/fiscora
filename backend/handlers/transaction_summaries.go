package handlers

import (
	"math"
	"net/http"
	"strconv"

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
		return DataBaseQueryError(c, err)
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
			return DataBaseQueryError(c, err)
		}

		monthInfo := types.GetMonthInfo(transactions, dateRange)

		yearInfo[i] = monthInfo
	}

	return c.JSON(http.StatusOK, yearInfo)
}

func (h *APIHandler) HandleGetTransactionsYearInfoPerType(c echo.Context) error {
	incomeParam := c.QueryParam("income")
	income, err := strconv.ParseBool(incomeParam)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid income type")
	}

	transactionTypes := make(map[string]float64)

	if income {
		for _, expenseType := range types.IncomeTypes {
			transactionTypes[expenseType] = 0
		}
	} else {
		for _, expenseType := range types.ExpenseTypes {
			transactionTypes[expenseType] = 0
		}
	}

	for i := 1; i < 13; i++ {
		transactionTypesMonth, err := getTransactionsPerType(c, h.DB)
		if err != nil {
			return err
		}

		for key := range transactionTypes {
			if val, ok := transactionTypesMonth[key]; ok {
				transactionTypes[key] += val
			} else {
				return InternalServerError(c, "Error getting yearly transaction info per type: invalid key")
			}
		}
	}

	for key := range transactionTypes {
		transactionTypes[key] /= 12
	}

	for key := range transactionTypes {
		if transactionTypes[key] == 0 {
			delete(transactionTypes, key)
		}
	}

	return c.JSON(http.StatusOK, transactionTypes)
}

func (h *APIHandler) HandleGetTransactionsPerType(c echo.Context) error {
	transactionTypes, err := getTransactionsPerType(c, h.DB)
	if err != nil {
		return err
	}

	for key := range transactionTypes {
		if transactionTypes[key] == 0 {
			delete(transactionTypes, key)
		}
	}

	return c.JSON(http.StatusOK, transactionTypes)
}

func getTransactionsPerType(c echo.Context, db *database.Queries) (map[string]float64, error) {
	userId := getUserId(c)
	month := getMonth(c)
	year := getYear(c)
	dateRange := getMonthRange(month, year)

	incomeParam := c.QueryParam("income")
	income, err := strconv.ParseBool(incomeParam)
	if err != nil {
		return nil, c.String(http.StatusBadRequest, "Invalid income type")
	}

	dbTransactions, err := getTransactionsFromDB(c.Request().Context(), c.QueryParam("income"), userId, dateRange, db)
	if err != nil {
		return nil, DataBaseQueryError(c, err)
	}

	transactions := types.ToTransactions(dbTransactions, dateRange)
	transactionTypes := make(map[string]float64)

	if income {
		for _, expenseType := range types.IncomeTypes {
			transactionTypes[expenseType] = 0
		}
	} else {
		for _, expenseType := range types.ExpenseTypes {
			transactionTypes[expenseType] = 0
		}
	}

	for _, transaction := range transactions {
		transactionTypes[transaction.Type] += math.Abs(transaction.Amount)
	}

	return transactionTypes, nil
}
