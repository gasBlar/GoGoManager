package routes

import (
	"net/http"

	"github.com/gasBlar/GoGoManager/api/v1/controllers"
)

func SetupExampleRoutes() {
	http.HandleFunc("/example", controllers.HelloHandler)
}
