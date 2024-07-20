package handlers

import "github.com/tvgelderen/budget-buddy/auth"

type APIHandler struct {
	// DB *database.Queries
    AuthService *auth.AuthService
}

func NewAPIHandler(auth *auth.AuthService) *APIHandler {
	return &APIHandler{
		// DB: database.New(db),
        AuthService: auth,
	}
}
