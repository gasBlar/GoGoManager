package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetAWSConfig returns AWS configuration
func GetAWSConfig() (region string) {
	region = os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-east-1" // Default to us-east-1 if not provided
	}

	return region
}
