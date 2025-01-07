package controllers

import (
	"fmt"
	"net/http"

	"github.com/gasBlar/GoGoManager/api/v1/services"
	"github.com/gasBlar/GoGoManager/utils"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	message := services.GetHelloWorldMessage()

	fmt.Fprintf(w, message)
}

func HelloMeHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*utils.Claims)

	utils.Response(w, http.StatusOK, "", user)
}
