package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Data interface{} `json:"data"`
}
type ResponseMessage struct {
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, statusCode int, message string, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	var response any

	if message != "" {
		response = &ResponseMessage{Message: message}
	} else if payload != nil {
		response = payload
	}

	res, _ := json.Marshal(response)
	w.Write(res)
}
