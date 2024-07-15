package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func HandleOAuthLogin(c echo.Context) error {
	if user, err := gothic.CompleteUserAuth(c.Response().Writer, c.Request()); err == nil {
		return c.JSON(http.StatusOK, user)
	}

	gothic.BeginAuthHandler(c.Response().Writer, c.Request())
	return nil
}

func HandleOAuthCallback(c echo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response().Writer, c.Request())
	if err != nil {
		return err
	}

    setToken(c.Response().Writer, user.IDToken)

	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/")
}

func setToken(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     "AccessToken",
		Value:    token,
		MaxAge:   36000,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}
