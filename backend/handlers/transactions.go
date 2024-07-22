package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tvgelderen/budget-buddy/models"
)

func (h *APIHandler) HandleGetTransactionIntervals(c echo.Context) error {
    intervals := make([]string, len(models.TransactionIntervals))

	for idx, interval := range models.TransactionIntervals {
		intervals[idx] = interval
	}

	return c.JSON(http.StatusOK, intervals)
}

func (h *APIHandler) HandleGetIncomeTypes(c echo.Context) error {
    incomeTypes := make([]string, len(models.IncomeTypes))

    for idx, incomeType := range models.IncomeTypes {
        incomeTypes[idx] = incomeType
    }

    return c.JSON(http.StatusOK, incomeTypes)
}

func (h *APIHandler) HandleGetExpenseTypes(c echo.Context) error {
    expenseTypes := make([]string, len(models.ExpenseTypes))

    for idx, expenseType := range models.ExpenseTypes {
        expenseTypes[idx] = expenseType
    }

    return c.JSON(http.StatusOK, expenseTypes)
}

func (h *APIHandler) HandlePostTransaction(c echo.Context) error {
    fmt.Println("handle post transaction")

    return c.NoContent(http.StatusOK)
}
