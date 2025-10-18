package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	OpenAIAPIKey string
}

var AppConfig *Config

// LoadConfig loads environment variables from .env file
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading from environment")
	}

	AppConfig = &Config{
		OpenAIAPIKey: getEnv("OPENAI_API_KEY", ""),
	}

	if AppConfig.OpenAIAPIKey == "" {
		log.Fatal("OPENAI_API_KEY is not set. Please check your .env file or environment variables.")
	}
}

// Helper function to get an environment variable or return a default
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
