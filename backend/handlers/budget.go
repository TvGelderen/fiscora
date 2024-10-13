package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tvgelderen/fiscora/database"
	"github.com/tvgelderen/fiscora/types"
)

func (h *APIHandler) HandleGetBudget(c echo.Context) error {
	userId := getUserId(c)

	budgets, err := getBudgetsFromDB(c.Request().Context(), userId, h.DB)
	if err != nil {
		return DataBaseQueryError(c, err)
	}

	return c.JSON(http.StatusOK, budgets)
}

func (h *APIHandler) HandleCreateBudget(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	budget := types.BudgetCreateRequest{}
	err := decoder.Decode(&budget)
	if err != nil {
		log.Errorf("Error decoding request body: %v", err.Error())
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	userId := getUserId(c)

	dbBudget, err := h.DB.CreateBudget(c.Request().Context(), database.CreateBudgetParams{
		ID:          generateRandomString(16),
		UserID:      userId,
		Name:        budget.Name,
		Description: budget.Description,
		Amount:      strconv.FormatFloat(budget.Amount, 'f', -1, 64),
		StartDate:   budget.StartDate,
		EndDate:     budget.EndDate,
	})
	if err != nil {
		return DataBaseQueryError(c, err)
	}

	return c.JSON(http.StatusCreated, types.ToBudget(dbBudget))
}

func (h *APIHandler) HandleUpdateBudget(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	budget := types.BudgetUpdateRequest{}
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

	dbBudget, err := h.DB.UpdateBudget(c.Request().Context(), database.UpdateBudgetParams{
		ID:          budgetId,
		UserID:      userId,
		Name:        budget.Name,
		Description: budget.Description,
		Amount:      strconv.FormatFloat(budget.Amount, 'f', -1, 64),
		StartDate:   budget.StartDate,
		EndDate:     budget.EndDate,
	})
	if err != nil {
		return DataBaseQueryError(c, err)
	}

	return c.JSON(http.StatusOK, types.ToBudget(dbBudget))
}

func (h *APIHandler) HandleDeleteBudget(c echo.Context) error {
	userId := getUserId(c)
	budgetId := c.Param("id")
	if budgetId == "" {
		log.Errorf("Error parsing budget id from request")
		return c.String(http.StatusBadRequest, "Error decoding request body")
	}

	err := h.DB.DeleteBudget(c.Request().Context(), database.DeleteBudgetParams{
		ID:     budgetId,
		UserID: userId,
	})
	if err != nil {
		return DataBaseQueryError(c, err)
	}

	return c.String(http.StatusOK, "Budget deleted successfully")
}
