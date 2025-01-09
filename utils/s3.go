package utils

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gasBlar/GoGoManager/config"
)

// UploadFileToS3 uploads a file to AWS S3 and returns the file URL
func UploadFileToS3(file multipart.File, filename string) (string, error) {
	bucketName := os.Getenv("AWS_BUCKET_NAME")
	region := os.Getenv("AWS_REGION")

	client, err := config.GetS3Client()
	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), filename)

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
		ACL:    types.ObjectCannedACLPublicRead,
	})

	if err != nil {
		log.Println("Failed to upload file to S3:", err)
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, key), nil
}
