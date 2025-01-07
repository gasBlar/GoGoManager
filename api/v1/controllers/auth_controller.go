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

	w.Header().Set("Content-Type", "application/json")

	if req.Type == "login" {
		res, err := services.Login(req)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Email is not found", http.StatusNotFound)
				return
			} else if err.Error() == "invalid password" {
				http.Error(w, "Invalid password", http.StatusUnauthorized)
				return
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)

	} else {
		res, err := services.Register(req)
		if err != nil {
			if err.Error() == "email already exists" {
				http.Error(w, "Email already exists", http.StatusConflict)
				return
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

func validateAuthRequest(req models.AuthLoginRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}
