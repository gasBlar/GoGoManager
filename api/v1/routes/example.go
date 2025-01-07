package routes

import (
	"github.com/gasBlar/GoGoManager/api/v1/controllers"
	"github.com/gorilla/mux"
)

func ExampleRoutes(r *mux.Router) {
	r.HandleFunc("/example", controllers.HelloHandler).Methods("GET")
}
