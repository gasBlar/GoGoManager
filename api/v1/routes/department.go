package routes

import (
	"database/sql"

	"github.com/gasBlar/GoGoManager/api/v1/controllers"
	"github.com/gasBlar/GoGoManager/api/v1/services"
	"github.com/gasBlar/GoGoManager/repository"
	"github.com/gorilla/mux"
)

func DepartmentRoutes(db *sql.DB, r *mux.Router) {
	// Initialize repository, service, and controller
	departmentRepo := repository.NewDepartmentRepository(db)
	departmentService := services.NewDepartmentService(departmentRepo)
	departmentController := controllers.NewDepartmentController(departmentService)

	// Register route
	r.HandleFunc("/department", departmentController.CreateDepartment).Methods("POST")
	r.HandleFunc("/department", controllers.GetDepartments(db)).Methods("GET")
}