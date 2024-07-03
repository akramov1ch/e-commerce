package config

import "os"

type Config struct {
    DBHost            string
    DBPort            string
    DBUser            string
    DBPassword        string
    DBName            string
    PORT              string
    USER_SERVICE_PORT string
}

func LoadConfig() Config {
    return Config{
        DBHost:            os.Getenv("DB_HOST"),
        DBPort:            os.Getenv("DB_PORT"),
        DBUser:            os.Getenv("DB_USER"),
        DBPassword:        os.Getenv("DB_PASSWORD"),
        DBName:            os.Getenv("DB_NAME"),
        PORT:              os.Getenv("PORT"),
        USER_SERVICE_PORT: os.Getenv("USER_SERVICE_PORT"),
    }
}
