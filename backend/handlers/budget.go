package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *APIHandler) HandleGetBudget(c echo.Context) error {
	userId := getUserId(c)

	budgets, err := getBudgetsFromDB(c.Request().Context(), userId, h.DB)
	if err != nil {
		return DataBaseQueryError(c, err)
	}

	return c.JSON(http.StatusOK, budgets)
}
