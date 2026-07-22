package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DatabaseURL    string
	JWTSecret      string
	JWTExpireHours int
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Note: .env file not found, using default environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("FATAL: JWT_SECRET environment variable is required and cannot be empty")
	}

	jwtExpireHours := 24
	if expireStr := os.Getenv("JWT_EXPIRE_HOURS"); expireStr != "" {
		if val, err := strconv.Atoi(expireStr); err == nil && val > 0 {
			jwtExpireHours = val
		}
	}

	return &Config{
		Port:           port,
		DatabaseURL:    dbURL,
		JWTSecret:      jwtSecret,
		JWTExpireHours: jwtExpireHours,
	}
}
