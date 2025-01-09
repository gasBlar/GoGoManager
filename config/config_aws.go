package config

import (
	"context"
	"log"
	"os"

	// "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Initialize AWS S3 client
func GetS3Client() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		log.Println("Failed to load AWS config:", err)
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}
