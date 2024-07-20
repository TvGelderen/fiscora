package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tvgelderen/budget-buddy/auth"
	"github.com/tvgelderen/budget-buddy/config"
)

func (h *APIHandler) HandleOAuthLogin(c echo.Context) error {
    url := h.AuthService.GoogleConfig.AuthCodeURL(config.Envs.SessionSecret)
    return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *APIHandler) HandleOAuthCallback(c echo.Context) error {
    state := c.Request().URL.Query().Get("state")
    if state != config.Envs.SessionSecret {
        return c.String(http.StatusForbidden, "Invalid state")
    }

    code := c.Request().URL.Query().Get("code")
    token, err := h.AuthService.GoogleConfig.Exchange(context.Background(), code)
    if err != nil {
        return c.String(http.StatusInternalServerError, "Error getting user token")
    }

    client := h.AuthService.GoogleConfig.Client(context.Background(), token)
    response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        return c.String(http.StatusInternalServerError, "Error getting user info")
    }

    var user auth.GoogleUser
    if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
        return c.String(http.StatusInternalServerError, "Error decoding user info")
    }

    setAccessToken(c.Response().Writer, token.AccessToken)
    
    return c.Redirect(http.StatusTemporaryRedirect, config.Envs.FrontendUrl)
}

func setAccessToken(w http.ResponseWriter, token string) {
    cookie := http.Cookie{
        Name: "AccessToken",
        Value: token,
        MaxAge: 60 * 60 * 24,
        Path: "/",
        HttpOnly: true,
        Secure: config.Envs.IsProduction,
    }

    http.SetCookie(w, &cookie)
}
