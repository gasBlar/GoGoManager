package routes

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	ExampleRoutes(s)
	AuthRoutes(s)

	return r
}
