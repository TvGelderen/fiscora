package handlers

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tvgelderen/budget-buddy/auth"
	"github.com/tvgelderen/budget-buddy/database"
)

type APIHandler struct {
	DB          *database.Queries
	AuthService *auth.AuthService
}

func NewAPIHandler(connection *sql.DB, auth *auth.AuthService) *APIHandler {
	return &APIHandler{
		DB:          database.New(connection),
		AuthService: auth,
	}
}

func InternalServerError(c echo.Context, err string) error {
    log.Error(err)
	return c.String(http.StatusInternalServerError, "Something went wrong")
}

func GetUserId(c echo.Context) uuid.UUID {
    return c.Get(userIdKey).(uuid.UUID)
}
