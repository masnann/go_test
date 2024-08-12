package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver  string
	DBName    string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	SSLMode   string
	JWTSecret string
}

var AppConfig Config

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return
	}

	AppConfig = Config{
		DBDriver:  GetEnv("DB_DRIVER"),
		DBName:    GetEnv("DB_NAME"),
		DBHost:    GetEnv("DB_HOST"),
		DBPort:    GetEnv("DB_PORT"),
		DBUser:    GetEnv("DB_USER"),
		DBPass:    GetEnv("DB_PASS"),
		SSLMode:   GetEnv("SSL_MODE"),
		JWTSecret: GetEnv("JWT_SECRET"),
	}
}

func GetEnv(key string, defaultValue ...string) string {
	value := os.Getenv(key)
	if value == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}
