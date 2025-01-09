package routes

import (
	"database/sql"

	"github.com/gasBlar/GoGoManager/api/v1/controllers"
	"github.com/gasBlar/GoGoManager/api/v1/services"
	"github.com/gasBlar/GoGoManager/repository"
	"github.com/gorilla/mux"
)

func EmployeeRoutes(db *sql.DB, r *mux.Router) {
	// Initialize repository, service, and controller
	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeController := controllers.NewEmployeeController(employeeService)

	// Register route
	r.HandleFunc("/employee", employeeController.CreateEmployee).Methods("POST")
	r.HandleFunc("/employee", employeeController.GetAllEmployees).Methods("GET")
	r.HandleFunc("/employee/{identityNumber}", employeeController.DeleteEmployee).Methods("DELETE")
	r.HandleFunc("/employee/{identityNumber}", employeeController.PatchEmployee).Methods("PATCH")
}
