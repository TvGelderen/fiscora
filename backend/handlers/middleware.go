package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tvgelderen/fiscora/auth"
)

const userIdKey = "user_id"

func (h *APIHandler) AuthorizeEndpoint(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := auth.GetId(c.Request())
		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		_, err = h.UserRepository.GetById(c.Request().Context(), id)
		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		c.Set(userIdKey, id)

		return next(c)
	}
}
