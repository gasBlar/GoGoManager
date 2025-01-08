package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	InitEnv()
}

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s not found", key)
	}
	return value
}
