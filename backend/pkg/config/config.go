package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

// LoadConfig reads configuration from .env file and environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists, but don't fail if it doesn't (production might use real env vars)
	_ = godotenv.Load()

	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}

	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBName == "" || cfg.JWTSecret == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return cfg, nil
}
