package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/tvgelderen/budget-buddy/handlers"

	_ "github.com/lib/pq"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    goth.UseProviders(google.New(os.Getenv("GOOGLE_ID"), os.Getenv("GOOGLE_SECRET"), "http://localhost:8080/api/oauth2/callback/google"))

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
		fmt.Println("PORT is missing, defaulting to 8080")
	}

    e := echo.New()
    
    e.GET("/api/auth", handlers.HandleOAuthLogin)
    e.GET("/api/oauth2/callback/:provider", handlers.OAuthCallback)

    e.Logger.Fatal(e.Start(port))
}
