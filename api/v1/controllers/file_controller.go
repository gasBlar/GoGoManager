package controllers

import (
	"github.com/gasBlar/GoGoManager/api/v1/services"
	"github.com/gasBlar/GoGoManager/utils"
	"fmt"
	"net/http"
)

// UploadFile handles file upload request
func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Validate method
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form data
	err := r.ParseMultipartForm(10 << 20) // Max 10MB
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Get file from form
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File not found", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file format
	if !utils.IsValidFileFormat(handler.Filename) {
		http.Error(w, "Invalid file format (must be jpeg, jpg, png)", http.StatusBadRequest)
		return
	}

	// Validate file size
	if handler.Size > 100*1024 {
		http.Error(w, "File size must be max 100 KiB", http.StatusBadRequest)
		return
	}

	// Call service to handle file upload
	fileURL, err := services.HandleFileUpload(file, handler.Filename)
	if err != nil {
		http.Error(w, "Failed to upload to S3", http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"uri": "%s"}`, fileURL)))
}
