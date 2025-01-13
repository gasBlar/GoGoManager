package utils

import (
	"net/url"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Fungsi untuk validasi bahwa file URI adalah gambar
func IsImageFileURI(fl validator.FieldLevel) bool {
	uri := fl.Field().String()
	// Parse URI
	parsedURI, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}

	// Ambil ekstensi file
	ext := strings.ToLower(filepath.Ext(parsedURI.Path))
	// Daftar ekstensi file gambar yang valid
	validExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	// Cek apakah ekstensi termasuk valid
	return validExtensions[ext]
}

func ValidateAddEmployee(employee interface{}) error {
	_ = validate.RegisterValidation("imagefileuri", IsImageFileURI)
	return validate.Struct(employee)
}
