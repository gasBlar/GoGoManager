package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gasBlar/GoGoManager/api/v1/services"
	"github.com/gasBlar/GoGoManager/models"
	"github.com/go-playground/validator/v10"
)

func LoginRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req models.AuthLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = validateAuthRequest(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Validation failed: %v", err), http.StatusBadRequest)
		return
	}

	// response := map[string]string{
	// 	"message": fmt.Sprintf("Successfully received request: %s", req),
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if req.Type == "login" {
		res, err := services.Login(req)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Email is not found", http.StatusNotFound)
				return
			} else {
				http.Error(w, fmt.Sprintf("Error logging in: %v", err), http.StatusInternalServerError)
				return
			}
		}

		json.NewEncoder(w).Encode(res)

	} else {
		res, err := services.Register(req)
		if err != nil {
			if err.Error() == "email already exists" {
				http.Error(w, "Email already exists", http.StatusConflict)
				return
			} else {
				http.Error(w, fmt.Sprintf("Error registering: %v", err), http.StatusInternalServerError)
			}
		}

		json.NewEncoder(w).Encode(res)
	}

	json.NewEncoder(w).Encode("asfas")

}

func validateAuthRequest(req models.AuthLoginRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}
