package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tvgelderen/fiscora/auth"
	"github.com/tvgelderen/fiscora/config"
	"github.com/tvgelderen/fiscora/database"
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

	error := c.Request().URL.Query().Get("error")
	if error != "" {
		return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/login", config.Envs.FrontendUrl))
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
		if !database.NoRowsFound(err) {
			return InternalServerError(c, fmt.Sprintf("Error getting user from db: %v", err.Error()))
		}

		user, err = h.DB.CreateUser(c.Request().Context(), database.CreateUserParams{
			ID:         uuid.New(),
			Provider:   "google",
			ProviderID: googleUser.Id,
			Username:   googleUser.Name,
			Email:      googleUser.Email,
		})
		if err != nil {
			return InternalServerError(c, fmt.Sprintf("Error creating user: %v", err.Error()))
		}
	}

	authToken, err := auth.CreateToken(user.ID, user.Username, user.Email)
	if err != nil {
		return InternalServerError(c, fmt.Sprintf("Error creating auth token: %v", err.Error()))
	}

	auth.SetToken(c.Response().Writer, authToken)

	return c.Redirect(http.StatusTemporaryRedirect, config.Envs.FrontendUrl)
}

func (h *APIHandler) HandleDemoLogin(c echo.Context) error {
	demo, err := h.DB.GetUserByEmail(c.Request().Context(), "demo")
	if err != nil {
		return InternalServerError(c, fmt.Sprintf("Error getting demo user from db: %v", err.Error()))
	}

	authToken, err := auth.CreateToken(demo.ID, demo.Username, demo.Email)
	if err != nil {
		return InternalServerError(c, fmt.Sprintf("Error creating auth token for demo user: %v", err.Error()))
	}

	auth.SetToken(c.Response().Writer, authToken)

	return c.Redirect(http.StatusTemporaryRedirect, config.Envs.FrontendUrl)
}

func (h *APIHandler) HandleLogout(c echo.Context) error {
	auth.DeleteToken(c.Response().Writer)

	return c.Redirect(http.StatusTemporaryRedirect, config.Envs.FrontendUrl)
}
