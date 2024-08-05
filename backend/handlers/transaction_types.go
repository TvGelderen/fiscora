package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
