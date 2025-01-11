package services

import (
	"github.com/gasBlar/GoGoManager/utils"
	"mime/multipart"
)

// HandleFileUpload handles file validation and uploads to S3
func HandleFileUpload(file multipart.File, filename string) (string, error) {
	// Upload file to S3
	fileURL, err := utils.UploadFileToS3(file, filename)
	if err != nil {
		return "", err
	}

	return fileURL, nil
}