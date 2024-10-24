package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tvgelderen/fiscora/repository"
	"github.com/tvgelderen/fiscora/types"
)

func (h *APIHandler) HandleGetBudgets(c echo.Context) error {
	userId := getUserId(c)
	budgets, err := h.BudgetRepository.Get(c.Request().Context(), userId)
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Errorf("Error getting budgets from db: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	returnBudgets := make([]types.BudgetReturn, len(*budgets))
	for idx, budget := range *budgets {
		returnBudgets[idx] = types.ToBudgetReturn(&budget)
	}

	return c.JSON(http.StatusOK, returnBudgets)
}

func (h *APIHandler) HandleGetBudget(c echo.Context) error {
	userId := getUserId(c)
	budgetId := c.Param("id")
	if budgetId == "" {
		log.Errorf("Error parsing budget id from request")
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	budget, err := h.BudgetRepository.GetById(c.Request().Context(), userId, budgetId)
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Errorf("Error getting budget from db: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}
	if budget.UserID != userId {
		return c.NoContent(http.StatusForbidden)
	}

	transactions, err := h.TransactionRepository.GetByBudgetId(c.Request().Context(), userId, budgetId)
	if err != nil {
		log.Errorf("Error getting budget transactions from db: %v", err.Error())
	}

	returnBudget := types.ToBudgetReturn(budget)

	if len(*transactions) != 0 {
		returnBudget.Transactions = types.ToTransactionReturns(transactions)
	}

	return c.JSON(http.StatusOK, returnBudget)
}

func (h *APIHandler) HandleCreateBudget(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	budgetForm := types.BudgetForm{}
	err := decoder.Decode(&budgetForm)
	if err != nil {
		log.Errorf("Error decoding request body: %v", err.Error())
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	userId := getUserId(c)
	budgetId := generateRandomString(16)

	budget, err := h.BudgetRepository.Add(c.Request().Context(), repository.CreateBudgetParams{
		ID:          budgetId,
		UserID:      userId,
		Name:        budgetForm.Name,
		Description: budgetForm.Description,
		Amount:      strconv.FormatFloat(budgetForm.Amount, 'f', -1, 64),
		StartDate:   budgetForm.StartDate,
		EndDate:     budgetForm.EndDate,
	})
	if err != nil {
		log.Errorf("Error creating budget: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	expenses := make([]repository.BudgetExpense, len(budgetForm.Expenses))
	for idx, expense := range budgetForm.Expenses {
		budgetExpense, err := h.BudgetRepository.AddExpense(c.Request().Context(), repository.CreateBudgetExpenseParams{
			BudgetID:        budgetId,
			Name:            expense.Name,
			AllocatedAmount: strconv.FormatFloat(expense.AllocatedAmount, 'f', -1, 64),
		})
		if err != nil {
			log.Errorf("Error creating budget expense: %v", err.Error())
			return c.String(http.StatusInternalServerError, "Something went wrong")
		}

		expenses[idx] = *budgetExpense
	}

	returnBudget := types.ToBudgetReturn(&repository.BudgetWithExpenses{
		Budget:   *budget,
		Expenses: expenses,
	})

	return c.JSON(http.StatusCreated, returnBudget)
}

func (h *APIHandler) HandleUpdateBudget(c echo.Context) error {
	userId := getUserId(c)
	budgetId := c.Param("id")
	if budgetId == "" {
		log.Errorf("Error parsing budget id from request")
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	decoder := json.NewDecoder(c.Request().Body)
	budgetForm := types.BudgetForm{}
	err := decoder.Decode(&budgetForm)
	if err != nil {
		log.Errorf("Error decoding request body: %v", err.Error())
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	err = h.BudgetRepository.Update(c.Request().Context(), repository.UpdateBudgetParams{
		ID:          budgetId,
		UserID:      userId,
		Name:        budgetForm.Name,
		Description: budgetForm.Description,
		Amount:      strconv.FormatFloat(budgetForm.Amount, 'f', -1, 64),
		StartDate:   budgetForm.StartDate,
		EndDate:     budgetForm.EndDate,
	})
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Errorf("Error updating budget: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	budgetExpenses, err := h.BudgetRepository.GetExpenses(c.Request().Context(), budgetId)
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Errorf("Error updating budget expense: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	for _, expense := range budgetForm.Expenses {
		if expense.ID == -1 {
			_, err := h.BudgetRepository.AddExpense(c.Request().Context(), repository.CreateBudgetExpenseParams{
				BudgetID:        budgetId,
				Name:            expense.Name,
				AllocatedAmount: strconv.FormatFloat(expense.AllocatedAmount, 'f', -1, 64),
			})
			if err != nil {
				log.Errorf("Error creating budget expense: %v", err.Error())
				return c.String(http.StatusInternalServerError, "Something went wrong")
			}
			continue
		}

		for _, budgetExpense := range *budgetExpenses {
			allocatedAmount := strconv.FormatFloat(expense.AllocatedAmount, 'f', -1, 64)
			if expense.ID == budgetExpense.ID &&
				(expense.Name != budgetExpense.Name || allocatedAmount != budgetExpense.AllocatedAmount) {
				err := h.BudgetRepository.UpdateExpense(c.Request().Context(), repository.UpdateBudgetExpenseParams{
					ID:              expense.ID,
					Name:            expense.Name,
					AllocatedAmount: allocatedAmount,
				})
				if err != nil {
					log.Errorf("Error updating budget expense: %v", err.Error())
					return c.String(http.StatusInternalServerError, "Something went wrong")
				}
			}
		}
	}

	budget, err := h.BudgetRepository.GetById(c.Request().Context(), userId, budgetId)

	returnBudget := types.ToBudgetReturn(budget)

	return c.JSON(http.StatusCreated, returnBudget)
}

func (h *APIHandler) HandleAddBudgetTransactions(c echo.Context) error {
	userId := getUserId(c)
	budgetId := c.Param("id")
	if budgetId == "" {
		log.Errorf("Error parsing budget id from request")
		return c.String(http.StatusBadRequest, "Invalid url parameter")
	}
	budgetExpenseId, err := strconv.ParseInt(c.Param("expense_id"), 10, 32)
	if err != nil {
		log.Errorf("Error parsing budget expense id from request")
		return c.String(http.StatusBadRequest, "Invalid url parameter")
	}
	if budgetExpenseId == -1 {
		return c.NoContent(http.StatusOK)
	}

	decoder := json.NewDecoder(c.Request().Body)
	var transactionIds []int32
	err = decoder.Decode(&transactionIds)
	if err != nil {
		log.Errorf("Error decoding request body: %v", err.Error())
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	for _, transactionId := range transactionIds {
		err := h.TransactionRepository.UpdateBudgetId(c.Request().Context(), repository.UpdateTransactionBudgetIdParams{
			UserID:          userId,
			ID:              transactionId,
			BudgetID:        budgetId,
			BudgetExpenseID: int32(budgetExpenseId),
		})
		if err != nil {
			if repository.NoRowsFound(err) {
				return c.NoContent(http.StatusNotFound)
			}
			log.Errorf("Error updating budget id for transaction: %v", err.Error())
			return c.String(http.StatusInternalServerError, "Something went wrong")
		}
	}

	transactions, err := h.TransactionRepository.GetByBudgetId(c.Request().Context(), userId, budgetId)
	if err != nil {
		log.Errorf("Error getting transactions: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, types.ToTransactionReturns(transactions))
}

func (h *APIHandler) HandleDeleteBudget(c echo.Context) error {
	userId := getUserId(c)
	budgetId := c.Param("id")
	if budgetId == "" {
		log.Errorf("Error parsing budget id from request")
		return c.String(http.StatusBadRequest, "Invalid url parameter")
	}

	err := h.BudgetRepository.Remove(c.Request().Context(), userId, budgetId)
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Errorf("Error deleting budget: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	return c.String(http.StatusOK, "Budget deleted successfully")
}

func (h *APIHandler) HandleDeleteBudgetExpense(c echo.Context) error {
	userId := getUserId(c)
	budgetId := c.Param("id")
	if budgetId == "" {
		log.Errorf("Error parsing budget id from request")
		return c.String(http.StatusBadRequest, "Invalid url parameter")
	}
	budgetExpenseId, err := strconv.ParseInt(c.Param("expense_id"), 10, 32)
	if err != nil {
		log.Errorf("Error parsing budget expense id from request")
		return c.String(http.StatusBadRequest, "Invalid url parameter")
	}
	if budgetExpenseId == -1 {
		return c.NoContent(http.StatusOK)
	}

	dbBudget, err := h.BudgetRepository.GetById(c.Request().Context(), userId, budgetId)
	if dbBudget.UserID != userId {
		return c.NoContent(http.StatusForbidden)
	}

	err = h.BudgetRepository.RemoveExpense(c.Request().Context(), int32(budgetExpenseId), budgetId)
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Errorf("Error deleting budget expense: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	return c.String(http.StatusOK, "Budget deleted successfully")
}
