package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gasBlar/GoGoManager/api/v1/services"
	"github.com/gasBlar/GoGoManager/models"
	"github.com/gasBlar/GoGoManager/utils"
	"github.com/go-playground/validator/v10"
)

func LoginRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req models.AuthLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, "Invalid request body", nil)
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
				utils.Response(w, http.StatusNotFound, "Email is not found", nil)
				return
			} else if err.Error() == "invalid password" {
				utils.Response(w, http.StatusInternalServerError, "Invalid password", nil)
				return
			} else {
				utils.Response(w, http.StatusInternalServerError, err.Error(), nil)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)

	} else {
		res, err := services.Register(req)
		if err != nil {
			if err.Error() == "email already exists" {
				utils.Response(w, http.StatusConflict, "Email already exists", nil)
				return
			} else {
				utils.Response(w, http.StatusInternalServerError, err.Error(), nil)
				return
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
