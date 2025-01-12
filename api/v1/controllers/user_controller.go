package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

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
	var req models.ProfileManagerUpdateRequest

	var rawBody map[string]json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&rawBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for key := range rawBody {
		if !utils.IsValidField(reflect.TypeOf(req), key) {
			http.Error(w, fmt.Sprintf("Unknown field: %s", key), http.StatusBadRequest)
			return
		}
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
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

func validateUpdateRequest(req models.ProfileManagerUpdateRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}
