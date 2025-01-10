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
)

var (
	awsRegion     string
	awsAccessKey  string
	awsSecretKey  string
	awsBucketName string
)

func init() {
	awsRegion = os.Getenv("AWS_REGION")
	awsAccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")

	awsBucketName = os.Getenv("AWS_S3_BUCKET_NAME")
}

// CreateSession creates a new AWS session
func CreateSession() (*session.Session, error) {
	if awsRegion == "" {
		awsRegion = "us-east-1"
	}

	return session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(
			awsAccessKey,
			awsSecretKey,
			"",
		),
	})
}

// UploadFileToS3 uploads a file to AWS S3 and returns the file URL
func UploadFileToS3(file multipart.File, filename string) (string, error) {
	log.Println("AWS_ACCESS_KEY_ID:", os.Getenv("AWS_ACCESS_KEY_ID"))
	log.Println("AWS_SECRET_ACCESS_KEY:", os.Getenv("AWS_SECRET_ACCESS_KEY"))

	log.Println(awsRegion + "|" + awsAccessKey + "|" + awsSecretKey + "|" + awsBucketName + "|")
	sess, err := CreateSession()
	if err != nil {
		return "Error create session: ", err
	}
	if awsRegion == "" {
		log.Fatal("Error: AWS_REGION environment variable is not set")
	}

	svc := s3.New(sess)

	key := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), filename)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(awsBucketName),
		Key:    aws.String(key),
		Body:   file,
	})

	if err != nil {
		log.Println("Failed to upload file to S3:", err)
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", awsBucketName, awsRegion, key), nil
}
