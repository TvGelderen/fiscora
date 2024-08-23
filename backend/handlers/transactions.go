package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tvgelderen/budget-buddy/database"
	"github.com/tvgelderen/budget-buddy/types"
)

func (h *APIHandler) HandleGetTransactions(c echo.Context) error {
	userId := getUserId(c)
	month := getMonth(c)
	year := getYear(c)
	dateRange := getMonthRange(month, year)

	transactions, err := getTransactionsFromDB(c.Request().Context(), c.QueryParam("income"), userId, dateRange, h.DB)
	if err != nil {
		return DataBaseQueryError(c, err)
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

	userId := getUserId(c)

	_, err = h.DB.CreateTransaction(c.Request().Context(), database.CreateTransactionParams{
		UserID:       userId,
		Amount:       strconv.FormatFloat(transaction.Amount, 'f', -1, 64),
		Description:  transaction.Description,
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
		return DataBaseQueryError(c, err)
	}

	return c.String(http.StatusCreated, "Transaction created successfully")
}

func (h *APIHandler) HandleUpdateTransaction(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	transaction := types.TransactionUpdateRequest{}
	err := decoder.Decode(&transaction)
	if err != nil {
		log.Errorf("Error decoding request body: %v", err.Error())
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	userId := getUserId(c)
	transactionIdParam := c.Param("id")
	transactionId, err := strconv.ParseInt(transactionIdParam, 10, 64)
	if err != nil {
		log.Errorf("Error parsing transaction id from request: %v", err.Error())
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	err = h.DB.UpdateTransaction(c.Request().Context(), database.UpdateTransactionParams{
		ID:           transactionId,
		UserID:       userId,
		Amount:       strconv.FormatFloat(transaction.Amount, 'f', -1, 64),
		Description:  transaction.Description,
		Type:         transaction.Type,
		Recurring:    transaction.Recurring,
		StartDate:    transaction.StartDate,
		EndDate:      transaction.EndDate,
		Interval:     transaction.Interval.NullString,
		DaysInterval: transaction.DaysInterval.NullInt32,
		Updated:      time.Now().UTC(),
	})
	if err != nil {
		return DataBaseQueryError(c, err)
	}

	return c.NoContent(204)
}

func (h *APIHandler) HandleDeleteTransaction(c echo.Context) error {
	userId := getUserId(c)
	transactionIdParam := c.Param("id")

	transactionId, err := strconv.ParseInt(transactionIdParam, 10, 64)
	if err != nil {
		log.Errorf("Error parsing transaction id from request: %v", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err = h.DB.DeleteTransaction(c.Request().Context(), database.DeleteTransactionParams{
		UserID: userId,
		ID:     transactionId,
	})
	if err != nil {
		return DataBaseQueryError(c, err)
	}

	return c.NoContent(204)
}
