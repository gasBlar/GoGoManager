package controllers

import (
	"database/sql"
	"encoding/json"
	"log"

	"net/http"

	"github.com/gasBlar/GoGoManager/api/v1/services"
	"github.com/gasBlar/GoGoManager/models"
	"github.com/gasBlar/GoGoManager/utils"
)

type DepartmentController struct {
	Service *services.DepartmentService
}

func NewDepartmentController(service *services.DepartmentService) *DepartmentController {
	return &DepartmentController{Service: service}
}

func (c *DepartmentController) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*utils.Claims)
	var department models.Department
	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.Service.CreateDepartment(&department, user.Id); err != nil {
		log.Println("Error creating department:", err)
		http.Error(w, "Error creating department", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Department created successfully"})
}

func GetDepartments(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		departments, err := services.GetAllDepartments(r.Context(), db)
		if err != nil {
			http.Error(w, "Failed to retrieve departments", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(departments); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}