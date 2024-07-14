package handlers

import (
	"database/sql"
)

type APIHandler struct {
	// DB *database.Queries
}

func NewAPIHandler(db *sql.DB) *APIHandler {
	return &APIHandler{
		// DB: database.New(db),
	}
}
