package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	IsProduction       bool
	PublicHost         string
	Port               string
	FrontendUrl        string
	DBConnectionString string
	SessionSecret      string
	GoogleID           string
	GoogleSecret       string
	GoogleCallback     string
}

var Envs = getEnvironment()

func getEnvironment() Environment {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Environment{
		IsProduction:       getBoolEnv("IS_PRODUCTION", false),
		PublicHost:         getEnv("PUBLIC_HOST", "http://localhost"),
		Port:               getEnv("PORT", ":8080"),
		FrontendUrl:        getEnv("FRONTEND_URL", "http://localhost:5173"),
		DBConnectionString: getEnv("DB_CONNECTION_STRING", ""),
		SessionSecret:      getEnv("SESSION_SECRET", ""),
		GoogleID:           getEnv("GOOGLE_ID", ""),
		GoogleSecret:       getEnv("GOOGLE_SECRET", ""),
		GoogleCallback:     getEnv("GOOGLE_CALLBACK", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
    log.Printf("%s not found in environment, defaulting to: %s", key, fallback)
	return fallback
}

func getBoolEnv(key string, fallback bool) bool {
	if value := getEnv(key, ""); value != "" {
		return value == "true"
	}
    log.Printf("%s not found in environment, defaulting to: %v", key, fallback)
	return fallback
}
