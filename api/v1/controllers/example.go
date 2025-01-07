package controllers

import (
	"fmt"
	"net/http"

	"github.com/gasBlar/GoGoManager/api/v1/services"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	message := services.GetHelloWorldMessage()

	fmt.Fprintf(w, message)
}
