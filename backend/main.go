package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/tvgelderen/budget-buddy/handlers"
)

func main() {
    godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
		fmt.Println("PORT is missing, defaulting to 8080")
	}

    e := echo.New()
    
    e.GET("/api/transactions", handlers.HandleGetTransaction)
    e.POST("/api/transactions", handlers.HandlePostTransaction)

    e.Logger.Fatal(e.Start(port))
}
