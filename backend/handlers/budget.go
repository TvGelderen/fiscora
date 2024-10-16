package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tvgelderen/fiscora/repository"
	"github.com/tvgelderen/fiscora/types"
)

func (h *APIHandler) HandleGetBudgets(c echo.Context) error {
	userId := getUserId(c)

	budgets, err := getBudgetsFromDB(c.Request().Context(), userId, h.DB)
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Error(fmt.Sprintf("Error getting budgets from db: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, budgets)
}

func (h *APIHandler) HandleGetBudget(c echo.Context) error {
	userId := getUserId(c)
	budgetId := c.Param("id")
	if budgetId == "" {
		log.Errorf("Error parsing budget id from request")
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	dbBudget, err := h.DB.GetBudget(c.Request().Context(), budgetId)
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Error(fmt.Sprintf("Error getting budget from db: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}
	if dbBudget.UserID != userId {
		return c.NoContent(http.StatusForbidden)
	}

	dbBudgetExpenses, err := h.DB.GetBudgetExpenses(c.Request().Context(), budgetId)
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Error(fmt.Sprintf("Error getting budget expenses from db: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	expenses := make([]types.BudgetExpenseReturn, len(dbBudgetExpenses))
	for idx, dbBudgetExpense := range dbBudgetExpenses {
		expenses[idx] = types.ToBudgetExpense(dbBudgetExpense)
	}

	budget := types.ToBudget(dbBudget)
	budget.Expenses = expenses

	return c.JSON(http.StatusOK, budget)
}

func (h *APIHandler) HandleCreateBudget(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	budget := types.BudgetForm{}
	err := decoder.Decode(&budget)
	if err != nil {
		log.Errorf("Error decoding request body: %v", err.Error())
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	userId := getUserId(c)
	budgetId := generateRandomString(16)

	dbBudget, err := h.DB.CreateBudget(c.Request().Context(), repository.CreateBudgetParams{
		ID:          budgetId,
		UserID:      userId,
		Name:        budget.Name,
		Description: budget.Description,
		Amount:      strconv.FormatFloat(budget.Amount, 'f', -1, 64),
		StartDate:   budget.StartDate,
		EndDate:     budget.EndDate,
	})
	if err != nil {
		log.Error(fmt.Sprintf("Error creating budget: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	expenses := make([]types.BudgetExpenseReturn, len(budget.Expenses))
	for idx, expense := range budget.Expenses {
		dbBudgetExpense, err := h.DB.CreateBudgetExpense(c.Request().Context(), repository.CreateBudgetExpenseParams{
			BudgetID:        budgetId,
			Name:            expense.Name,
			AllocatedAmount: strconv.FormatFloat(expense.AllocatedAmount, 'f', -1, 64),
		})
		if err != nil {
			log.Error(fmt.Sprintf("Error creating budget expense: %v", err.Error()))
			return c.String(http.StatusInternalServerError, "Something went wrong")
		}

		expenses[idx] = types.ToBudgetExpense(dbBudgetExpense)
	}

	returnBudget := types.ToBudget(dbBudget)
	returnBudget.Expenses = expenses

	return c.JSON(http.StatusCreated, returnBudget)
}

func (h *APIHandler) HandleUpdateBudget(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	budget := types.BudgetForm{}
	err := decoder.Decode(&budget)
	if err != nil {
		log.Errorf("Error decoding request body: %v", err.Error())
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	userId := getUserId(c)
	budgetId := c.Param("id")
	if budgetId == "" {
		log.Errorf("Error parsing budget id from request")
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	dbBudget, err := h.DB.UpdateBudget(c.Request().Context(), repository.UpdateBudgetParams{
		ID:          budgetId,
		UserID:      userId,
		Name:        budget.Name,
		Description: budget.Description,
		Amount:      strconv.FormatFloat(budget.Amount, 'f', -1, 64),
		StartDate:   budget.StartDate,
		EndDate:     budget.EndDate,
	})
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Error(fmt.Sprintf("Error updating budget: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	dbBudgetExpenses, err := h.DB.GetBudgetExpenses(c.Request().Context(), budgetId)
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Error(fmt.Sprintf("Error updating budget expense: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	expenses := make([]types.BudgetExpenseReturn, len(budget.Expenses))
	for idx, expense := range budget.Expenses {
		if expense.ID == -1 {
			createdExpense, err := h.DB.CreateBudgetExpense(c.Request().Context(), repository.CreateBudgetExpenseParams{
				BudgetID:        budgetId,
				Name:            expense.Name,
				AllocatedAmount: strconv.FormatFloat(expense.AllocatedAmount, 'f', -1, 64),
			})
			if err != nil {
				log.Error(fmt.Sprintf("Error creating budget expense: %v", err.Error()))
				return c.String(http.StatusInternalServerError, "Something went wrong")
			}

			expenses[idx] = types.ToBudgetExpense(createdExpense)
			continue
		}

		for _, dbBudgetExpense := range dbBudgetExpenses {
			allocatedAmount := strconv.FormatFloat(expense.AllocatedAmount, 'f', -1, 64)
			currentAmount := strconv.FormatFloat(expense.AllocatedAmount, 'f', -1, 64)
			if expense.ID == dbBudgetExpense.ID &&
				(expense.Name != dbBudgetExpense.Name ||
					allocatedAmount != dbBudgetExpense.AllocatedAmount ||
					currentAmount != dbBudgetExpense.CurrentAmount) {
				updatedExpense, err := h.DB.UpdateBudgetExpense(c.Request().Context(), repository.UpdateBudgetExpenseParams{
					ID:              expense.ID,
					Name:            expense.Name,
					AllocatedAmount: allocatedAmount,
					CurrentAmount:   currentAmount,
				})
				if err != nil {
					log.Error(fmt.Sprintf("Error updating budget expense: %v", err.Error()))
					return c.String(http.StatusInternalServerError, "Something went wrong")
				}

				expenses[idx] = types.ToBudgetExpense(updatedExpense)
				break
			}
		}
	}

	returnBudget := types.ToBudget(dbBudget)
	returnBudget.Expenses = expenses

	return c.JSON(http.StatusCreated, returnBudget)
}

func (h *APIHandler) HandleDeleteBudget(c echo.Context) error {
	userId := getUserId(c)
	budgetId := c.Param("id")
	if budgetId == "" {
		log.Errorf("Error parsing budget id from request")
		return c.String(http.StatusBadRequest, "Invalid url parameter")
	}

	err := h.DB.DeleteBudget(c.Request().Context(), repository.DeleteBudgetParams{
		ID:     budgetId,
		UserID: userId,
	})
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Error(fmt.Sprintf("Error deleting budget: %v", err.Error()))
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
	budgetExpenseIdParam := c.Param("expense_id")
	budgetExpenseId, err := strconv.ParseInt(budgetExpenseIdParam, 10, 32)
	if err != nil {
		log.Errorf("Error parsing budget expense id from request")
		return c.String(http.StatusBadRequest, "Invalid url parameter")
	}
	if budgetExpenseId == -1 {
		return c.NoContent(http.StatusOK)
	}

	dbBudget, err := h.DB.GetBudget(c.Request().Context(), budgetId)
	if dbBudget.UserID != userId {
		return c.NoContent(http.StatusForbidden)
	}

	err = h.DB.DeleteBudgetExpense(c.Request().Context(), repository.DeleteBudgetExpenseParams{
		ID:       int32(budgetExpenseId),
		BudgetID: budgetId,
	})
	if err != nil {
		if repository.NoRowsFound(err) {
			return c.NoContent(http.StatusNotFound)
		}
		log.Error(fmt.Sprintf("Error deleting budget expense: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	return c.String(http.StatusOK, "Budget deleted successfully")
}
