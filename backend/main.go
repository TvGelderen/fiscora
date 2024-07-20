package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tvgelderen/budget-buddy/auth"
	"github.com/tvgelderen/budget-buddy/config"
	"github.com/tvgelderen/budget-buddy/handlers"

	_ "github.com/lib/pq"
)

type Test struct {
	Message string `json:"message"`
}

func main() {
	env := config.Envs

	authService := auth.NewAuthService()
	handler := handlers.NewAPIHandler(authService)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/auth/:provider", handler.HandleOAuthLogin)
	e.GET("/auth/callback/:provider", handler.HandleOAuthCallback)

	e.GET("/ping", func(c echo.Context) error {
		body, _ := json.Marshal(Test{Message: "Pong"})
		return c.JSON(http.StatusOK, body)
	})

	e.Logger.Fatal(e.Start(env.Port))
}
