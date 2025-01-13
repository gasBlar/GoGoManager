package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/gasBlar/GoGoManager/api/v1/services"
	"github.com/gasBlar/GoGoManager/models"
	"github.com/gasBlar/GoGoManager/utils"
	"github.com/go-playground/validator/v10"
)

func GetDataUserHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*utils.Claims)
	result, err := services.GetUserProfile(user.Id)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.Response(w, http.StatusOK, "", result)
}

func UpdateDataUserHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*utils.Claims)
	var rBody []byte
	var err error
	if rBody, err = io.ReadAll(r.Body); err != nil {
		http.Error(w, fmt.Sprintf("Failed to read body: %s", err), http.StatusInternalServerError)
		return
	}

	var rawBody map[string]json.RawMessage
	if err := json.Unmarshal(rBody, &rawBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for key := range rawBody {
		if !utils.IsValidField(reflect.TypeOf(models.ProfileManagerUpdateRequest{}), key) {
			http.Error(w, fmt.Sprintf("Unknown field: %s", key), http.StatusBadRequest)
			return
		}
	}

	var req models.ProfileManagerUpdateRequest
	if err := json.Unmarshal(rBody, &req); err != nil {
		slog.Error("Error decoding request body", "error", err)
		utils.Response(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = validateUpdateRequest(req)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, fmt.Sprintf("Validation failed: %v", err), nil)
		return
	}

	result, err := services.UpdateUserProfile(user.Id, req)
	if err != nil {
		if err.Error() == "email already exists" {
			utils.Response(w, http.StatusConflict, "Email already exists", nil)
			return
		}
		utils.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.Response(w, http.StatusOK, "", result)
}

func validateURI(fl validator.FieldLevel) bool {
	uri := fl.Field().String()

	parsedURL, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}

	if parsedURL.Path == "" || strings.HasSuffix(parsedURL.Path, "/") {
		return false
	}

	return true
}

func validateNonEmpty(fl validator.FieldLevel) bool {
	// Get the field value
	fieldValue := fl.Field().String()

	// Check if the field is an empty string or the value "null" (as a string)
	if fieldValue == "" || strings.ToLower(fieldValue) == "null" {
		return false // Fail validation if empty or "null"
	}
	return true
}

func validateUpdateRequest(req models.ProfileManagerUpdateRequest) error {
	validate := validator.New()
	validate.RegisterValidation("uri", validateURI)
	validate.RegisterValidation("nonempty", validateNonEmpty)

	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}
