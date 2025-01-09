package utils

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gasBlar/GoGoManager/config"
)

// CreateSession creates a new AWS session
func CreateSession() (*session.Session, error) {
	region := config.GetAWSConfig()

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	return session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKey,
			secretKey,
			"",
		),
	})
}

// UploadFileToS3 uploads a file to AWS S3 and returns the file URL
func UploadFileToS3(file multipart.File, filename string) (string, error) {
	bucketName := os.Getenv("AWS_BUCKET_NAME")
	region := os.Getenv("AWS_REGION")

	log.Println("AWS_ACCESS_KEY_ID:", os.Getenv("AWS_ACCESS_KEY_ID"))
	log.Println("AWS_SECRET_ACCESS_KEY:", os.Getenv("AWS_SECRET_ACCESS_KEY"))
	sess, err := CreateSession()
	if err != nil {
		return "Error create session: ", err
	}
	if region == "" {
		log.Fatal("Error: AWS_REGION environment variable is not set")
	}

	svc := s3.New(sess)

	key := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), filename)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	})

	if err != nil {
		log.Println("Failed to upload file to S3:", err)
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, key), nil
}
