package utils

import "strings"

// IsValidFileFormat checks file extension
func IsValidFileFormat(filename string) bool {
	allowedExtensions := []string{".jpeg", ".jpg", ".png"}
	for _, ext := range allowedExtensions {
		if strings.HasSuffix(strings.ToLower(filename), ext) {
			return true
		}
	}
	return false
}
