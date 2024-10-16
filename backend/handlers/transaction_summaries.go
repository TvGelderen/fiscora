package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tvgelderen/fiscora/repository"
	"github.com/tvgelderen/fiscora/types"
)

func (h *APIHandler) HandleGetTransactionMonthInfo(c echo.Context) error {
	userId := getUserId(c)
	month := getMonth(c)
	year := getYear(c)

	dateRange := getMonthRange(month, year)
	dbTransactions, err := h.DB.GetBaseTransactionsBetweenDates(c.Request().Context(), repository.GetBaseTransactionsBetweenDatesParams{
		UserID:    userId,
		StartDate: dateRange.Start,
		EndDate:   dateRange.End,
		Limit:     repository.MaxFetchLimit,
		Offset:    0,
	})
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Error(fmt.Sprintf("Error getting transactions from db: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	monthInfo := types.GetMonthInfo(dbTransactions)

	return c.JSON(http.StatusOK, monthInfo)
}

func (h *APIHandler) HandleGetTransactionYearInfo(c echo.Context) error {
	userId := getUserId(c)
	year := getYear(c)

	yearInfo := make(map[int]types.MonthInfoReturn)

	for month := 1; month < 13; month++ {
		dateRange := getMonthRange(month, year)
		dbTransactions, err := h.DB.GetBaseTransactionsBetweenDates(c.Request().Context(), repository.GetBaseTransactionsBetweenDatesParams{
			UserID:    userId,
			StartDate: dateRange.Start,
			EndDate:   dateRange.End,
			Limit:     repository.MaxFetchLimit,
			Offset:    0,
		})
		if err != nil {
			if repository.NoRowsFound(err) {
				return c.NoContent(http.StatusNotFound)
			}
			log.Error(fmt.Sprintf("Error getting transactions from db: %v", err.Error()))
			return c.String(http.StatusInternalServerError, "Something went wrong")
		}

		monthInfo := types.GetMonthInfo(dbTransactions)

		yearInfo[month] = monthInfo
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
				log.Error("Error getting yearly transaction info per type: invalid key")
				return c.String(http.StatusInternalServerError, "Something went wrong")
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

func getTransactionsPerType(c echo.Context, db *repository.Queries) (map[string]float64, error) {
	userId := getUserId(c)
	month := getMonth(c)
	year := getYear(c)
	dateRange := getMonthRange(month, year)

	incomeParam := c.QueryParam("income")
	income, err := strconv.ParseBool(incomeParam)
	if err != nil {
		return nil, c.String(http.StatusBadRequest, "Invalid income type")
	}

	transactions, err := getTransactionsFromDB(c.Request().Context(), c.QueryParam("income"), userId, dateRange, db)
	if err != nil {
		if repository.NoRowsFound(err) {
			return nil, c.NoContent(http.StatusNotFound)
		}
		log.Error(fmt.Sprintf("Error getting transactions from db: %v", err.Error()))
		return nil, c.String(http.StatusInternalServerError, "Something went wrong")
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

	for _, transaction := range transactions {
		transactionTypes[transaction.Type] += math.Abs(transaction.Amount)
	}

	return transactionTypes, nil
}
