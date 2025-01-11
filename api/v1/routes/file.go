package routes

import (
	"github.com/gasBlar/GoGoManager/api/v1/controllers"
	"github.com/gorilla/mux"
)

func RegisterFileRoutes(r *mux.Router) {
	r.HandleFunc("/file", controllers.UploadFile).Methods("POST")
}
