package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tvgelderen/budget-buddy/types"
)

func (h *APIHandler) HandleGetMe(c echo.Context) error {
    id := c.Get(userIdKey)
    if id == nil {
        return InternalServerError(c, "Unable to get user id from context")
    }

    user, err := h.DB.GetUserById(c.Request().Context(), id.(uuid.UUID))
    if err != nil {
        return InternalServerError(c, fmt.Sprintf("Error getting user from db: %v", err.Error()))
    }

    return c.JSON(http.StatusOK, types.ToUser(user))
}
