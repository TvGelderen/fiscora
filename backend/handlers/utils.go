package handlers

type APIHandler struct {
	// DB *database.Queries
}

func NewAPIHandler() *APIHandler {
	return &APIHandler{
		// DB: database.New(db),
	}
}
