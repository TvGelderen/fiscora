package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tvgelderen/budget-buddy/auth"
	"github.com/tvgelderen/budget-buddy/config"
	"github.com/tvgelderen/budget-buddy/database"
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
		return InternalServerError(c, fmt.Sprintf("Error getting user token: %v", err.Error()))
	}

	client := h.AuthService.GoogleConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return InternalServerError(c, fmt.Sprintf("Error getting user info: %v", err.Error()))
	}

	var googleUser auth.GoogleUser
	if err = json.NewDecoder(response.Body).Decode(&googleUser); err != nil {
		return InternalServerError(c, fmt.Sprintf("Error decoding user info: %v", err.Error()))
	}

	user, err := h.DB.GetUserByProviderId(c.Request().Context(), database.GetUserByProviderIdParams{
		Provider:   "google",
		ProviderID: googleUser.Id,
	})
	if err != nil {
		if !strings.Contains(err.Error(), "no rows in result set") {
			return InternalServerError(c, fmt.Sprintf("Error getting user from db: %v", err.Error()))
		}

		user, err = h.DB.CreateUser(c.Request().Context(), database.CreateUserParams{
			ID:         uuid.New(),
			Provider:   "google",
			ProviderID: googleUser.Id,
			Username:   googleUser.Name,
			Email:      googleUser.Email,
			Created:    time.Now().UTC(),
			Updated:    time.Now().UTC(),
		})
		if err != nil {
			return InternalServerError(c, fmt.Sprintf("Error creating user: %v", err.Error()))
		}
	}

	fmt.Println(user)

	setAccessToken(c.Response().Writer, token.AccessToken)

	return c.Redirect(http.StatusTemporaryRedirect, config.Envs.FrontendUrl)
}

func setAccessToken(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     "AccessToken",
		Value:    token,
		MaxAge:   60 * 60 * 24,
		Path:     "/",
		HttpOnly: true,
		Secure:   config.Envs.IsProduction,
	}

	http.SetCookie(w, &cookie)
}
