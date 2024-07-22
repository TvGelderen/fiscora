package main

import (
	"log"
	"database/sql"

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

	if env.DBConnectionString == "" {
		log.Fatal("No database connection string found")
	}

	connection, err := sql.Open("postgres", env.DBConnectionString)
	if err != nil {
		log.Fatalf("Error establishing database connection: %s", err.Error())
	}

	authService := auth.NewAuthService()
	handler := handlers.NewAPIHandler(connection, authService)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/auth/:provider", handler.HandleOAuthLogin)
	e.GET("/auth/callback/:provider", handler.HandleOAuthCallback)

    e.GET("/auth/me", handler.AuthorizeEndpoint(handler.HandleGetMe))

    e.GET("/transactions/intervals", handler.HandleGetTransactionIntervals)
    e.GET("/transactions/income-types", handler.HandleGetIncomeTypes)
    e.GET("/transactions/expense-types", handler.HandleGetExpenseTypes)

    e.POST("/transactions", handler.AuthorizeEndpoint(handler.HandleCreateTransaction))

	e.Logger.Fatal(e.Start(env.Port))
}
