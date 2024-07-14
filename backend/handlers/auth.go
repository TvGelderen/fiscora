package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func HandleOAuthLogin(c echo.Context) error {
	if user, err := gothic.CompleteUserAuth(c.Response().Writer, c.Request()); err == nil {
		return c.JSON(http.StatusOK, user)
	}

	gothic.BeginAuthHandler(c.Response().Writer, c.Request())
    fmt.Println("handleauth")
	return nil
}

func OAuthCallback(c echo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response().Writer, c.Request())
	if err != nil {
		return err
	}

    userBytes, err := json.Marshal(user)
    if err != nil {
        return err
    }
    c.Response().Header().Add("Content-Type", "application/json")
    c.Response().Write(userBytes)

    fmt.Println("Redirecting")

    return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173")
}
