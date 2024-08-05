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

    // Enable users to view the site without logging in
	e.GET("/auth/demo", handler.HandleDemoLogin)

	e.GET("/auth/:provider", handler.HandleOAuthLogin)
	e.GET("/auth/callback/:provider", handler.HandleOAuthCallback)

    e.GET("/transactions/types/intervals", handler.HandleGetTransactionIntervals)
    e.GET("/transactions/types/income", handler.HandleGetIncomeTypes)
    e.GET("/transactions/types/expense", handler.HandleGetExpenseTypes)

    // Authorized endpoints
    authorized := e.Group("", handler.AuthorizeEndpoint)

    authorized.GET("/users/me", handler.HandleGetMe)

    authorized.GET("/transactions", handler.HandleGetTransactions)
    authorized.POST("/transactions", handler.HandleCreateTransaction)
    authorized.PUT("/transactions/:id", handler.HandleUpdateTransaction)
    authorized.DELETE("/transactions/:id", handler.HandleDeleteTransaction)

    authorized.GET("/transactions/summary/month", handler.HandleGetTransactionMonthInfo)
    authorized.GET("/transactions/summary/year", handler.HandleGetTransactionYearInfo)
    authorized.GET("/transactions/summary/expenses-per-type", handler.HandleGetExpensePerType)

	e.Logger.Fatal(e.Start(env.Port))
}
