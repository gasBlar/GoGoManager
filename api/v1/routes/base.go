package routes

import (
	"github.com/gasBlar/GoGoManager/api/v1/middleware"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.JWTMiddleware)
	s := r.PathPrefix("/api/v1").Subrouter()

	ExampleRoutes(s)
	AuthRoutes(s)
	UserRoutes(s)
	RegisterFileRoutes(s)

	return r
}
