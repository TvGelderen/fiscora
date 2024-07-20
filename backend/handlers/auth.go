package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tvgelderen/budget-buddy/config"
)

func HandleAuthCallback(c echo.Context) error {
    fmt.Println("Auth callback")
    
    fmt.Println(c.Request().Body)

    return c.Redirect(http.StatusOK, config.Envs.FrontendUrl)
}
