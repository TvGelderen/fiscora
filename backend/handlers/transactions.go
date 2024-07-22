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

func (h *APIHandler) HandleCreateTransaction(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	transaction := types.TransactionCreateRequest{}
	err := decoder.Decode(&transaction)
	if err != nil {
		log.Errorf("Error decoding request body: %v", err.Error())
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	fmt.Println(transaction)

	// TODO: validate transaction object

	userId := GetUserId(c)

	record, err := h.DB.CreateTransaction(c.Request().Context(), database.CreateTransactionParams{
		UserID:       userId,
		Amount:       strconv.FormatFloat(transaction.Amount, 'f', -1, 64),
		Description:  transaction.Description,
		Incoming:     transaction.Incoming,
		Type:         transaction.Type,
		Recurring:    transaction.Recurring,
		StartDate:    transaction.StartDate,
		EndDate:      transaction.EndDate.NullTime,
		Interval:     transaction.Interval.NullString,
		DaysInterval: transaction.DaysInterval.NullInt32,
		Created:      time.Now().UTC(),
		Updated:      time.Now().UTC(),
	})
    if err != nil {
        return InternalServerError(c, fmt.Sprintf("Error creating transaction: %v", err.Error()))
    }

    fmt.Println("Created transaction record")

	return c.JSON(http.StatusOK, types.ToTransaction(record))
}
