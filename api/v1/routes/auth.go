package routes

import (
	"github.com/gasBlar/GoGoManager/api/v1/controllers"
	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth", controllers.LoginRegisterHandler).Methods("POST")
}
