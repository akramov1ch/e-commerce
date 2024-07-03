package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	UserServiceURL string
	JWTSecret      string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := &Config{
		UserServiceURL: os.Getenv("USER_SERVICE_URL"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
	}

	return config
}
