package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tvgelderen/fiscora/auth"
	"github.com/tvgelderen/fiscora/config"
	"github.com/tvgelderen/fiscora/repository"
)

func (h *APIHandler) HandleOAuthLogin(c echo.Context) error {
	url := h.AuthService.GoogleConfig.AuthCodeURL(config.Envs.SessionSecret)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *APIHandler) HandleOAuthCallback(c echo.Context) error {
	query := c.Request().URL.Query()
	state := query.Get("state")
	if state != config.Envs.SessionSecret {
		return c.String(http.StatusForbidden, "Invalid state")
	}

	error := query.Get("error")
	if error != "" {
		return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/login", config.Envs.FrontendUrl))
	}

	code := query.Get("code")
	token, err := h.AuthService.GoogleConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Error(fmt.Sprintf("Error getting user token: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	client := h.AuthService.GoogleConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Error(fmt.Sprintf("Error getting user info: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	var googleUser auth.GoogleUser
	if err = json.NewDecoder(response.Body).Decode(&googleUser); err != nil {
		log.Error(fmt.Sprintf("Error decoding user info: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	user, err := h.UserRepository.GetByProviderId(c.Request().Context(), "google", googleUser.Id)
	if err != nil {
		if !repository.NoRowsFound(err) {
			log.Error(fmt.Sprintf("Error getting user from db: %v", err.Error()))
			return c.String(http.StatusInternalServerError, "Something went wrong")
		}

		user, err = h.UserRepository.Add(c.Request().Context(), repository.CreateUserParams{
			ID:         uuid.New(),
			Provider:   "google",
			ProviderID: googleUser.Id,
			Username:   googleUser.Name,
			Email:      googleUser.Email,
		})
		if err != nil {
			log.Error(fmt.Sprintf("Error creating user: %v", err.Error()))
			return c.String(http.StatusInternalServerError, "Something went wrong")
		}
	}

	authToken, err := auth.CreateToken(user.ID, user.Username, user.Email)
	if err != nil {
		log.Error(fmt.Sprintf("Error creating auth token: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	auth.SetToken(c.Response().Writer, authToken)

	return c.Redirect(http.StatusTemporaryRedirect, config.Envs.FrontendUrl)
}

func (h *APIHandler) HandleDemoLogin(c echo.Context) error {
	demo, err := h.UserRepository.GetByEmail(c.Request().Context(), "demo")
	if err != nil {
		log.Error(fmt.Sprintf("Error getting demo user from db: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	authToken, err := auth.CreateToken(demo.ID, demo.Username, demo.Email)
	if err != nil {
		log.Error(fmt.Sprintf("Error creating auth token for demo user: %v", err.Error()))
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	auth.SetToken(c.Response().Writer, authToken)

	return c.Redirect(http.StatusTemporaryRedirect, config.Envs.FrontendUrl)
}

func (h *APIHandler) HandleLogout(c echo.Context) error {
	auth.DeleteToken(c.Response().Writer)

	return c.NoContent(http.StatusOK)
}
