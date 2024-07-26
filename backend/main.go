package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tvgelderen/budget-buddy/auth"
	"github.com/tvgelderen/budget-buddy/config"
	"github.com/tvgelderen/budget-buddy/handlers"
	"github.com/tvgelderen/budget-buddy/seed"

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

    seedFlag := flag.Bool("seed", false, "Set to true to seed the database")
    flag.Parse()
    if *seedFlag {
        seed.Seed(connection)
    }

	authService := auth.NewAuthService()
	handler := handlers.NewAPIHandler(connection, authService)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/auth/:provider", handler.HandleOAuthLogin)
	e.GET("/auth/callback/:provider", handler.HandleOAuthCallback)

    e.GET("/transactions/intervals", handler.HandleGetTransactionIntervals)
    e.GET("/transactions/income-types", handler.HandleGetIncomeTypes)
    e.GET("/transactions/expense-types", handler.HandleGetExpenseTypes)

    // Authorized endpoints
    authorized := e.Group("", handler.AuthorizeEndpoint)

    authorized.GET("/auth/me", handler.HandleGetMe)

    authorized.GET("/transactions", handler.HandleGetTransactions)
    authorized.POST("/transactions", handler.HandleCreateTransaction)
    authorized.GET("/transactions/month-info", handler.HandleGetTransactionMonthInfo)

	e.Logger.Fatal(e.Start(env.Port))
}
