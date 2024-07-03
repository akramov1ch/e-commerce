package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiGatewayPort string
	UserServicePort string
	ProductServicePort string
	OrderServicePort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := &Config{
		ApiGatewayPort: os.Getenv("API_GATEWAY_PORT"),
		UserServicePort: os.Getenv("USER_SERVICE_PORT"),
		ProductServicePort: os.Getenv("PRODUCT_SERVICE_PORT"),
		OrderServicePort: os.Getenv("ORDER_SERVICE_PORT"),
	}

	return config
}
