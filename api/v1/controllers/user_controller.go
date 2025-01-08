package controllers

import (
	"net/http"

	"github.com/gasBlar/GoGoManager/api/v1/services"
	"github.com/gasBlar/GoGoManager/utils"
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
