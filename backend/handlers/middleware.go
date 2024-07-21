package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tvgelderen/budget-buddy/auth"
)

const userIdKey = "user_id"

func (h *APIHandler) AuthorizeEndpoint(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, err := auth.GetId(c.Request())
        if err != nil {
            return c.NoContent(http.StatusUnauthorized)
        }

        val, err := h.DB.GetUserExists(c.Request().Context(), id)
        if val == 0 || err != nil {
            return c.NoContent(http.StatusUnauthorized)
        }
        
        c.Set(userIdKey, id)

        return next(c)
    }
}
