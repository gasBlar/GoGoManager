package routes

import (
	"github.com/gasBlar/GoGoManager/api/v1/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/user", controllers.GetDataUserHandler).Methods("GET")
	// r.HandleFunc("/user", controllers.LoginRegisterHandler).Methods("PATCH")
}
