package controllers

import (
	"encoding/json"
	"log"

	"net/http"

	"github.com/gasBlar/GoGoManager/api/v1/services"
	"github.com/gasBlar/GoGoManager/models"
	"github.com/gorilla/mux"
	// "github.com/gasBlar/GoGoManager/utils"
	// "github.com/go-playground/validator/v10"
)

type EmployeeController struct {
	Service *services.EmployeeService
}

func NewEmployeeController(service *services.EmployeeService) *EmployeeController {
	return &EmployeeController{Service: service}
}

func (c *EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.Service.CreateEmployee(&employee); err != nil {
		log.Println("Error creating employee:", err)
		http.Error(w, "Error creating employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee created successfully"})
}

func (c *EmployeeController) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan data seluruh employee dari service
	employees, err := c.Service.GetAllEmployees()
	if err != nil {
		http.Error(w, "Error retrieving employees", http.StatusInternalServerError)
		return
	}

	// Menyusun respons dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// func GetEmployees(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		employees, err := services.GetAllEmployees(r.Context(), db)
// 		if err != nil {
// 			http.Error(w, "Failed to retrieve employees", http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		if err := json.NewEncoder(w).Encode(employees); err != nil {
// 			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }

func (c *EmployeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["identityNumber"]

	if err := c.Service.DeleteEmployee(id); err != nil {
		http.Error(w, "Error deleting employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee deleted successfully"})
}

func (c *EmployeeController) PatchEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.EmployeePatch
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Ambil identityNumber dari URL
	vars := mux.Vars(r)
	identityNumber := vars["identityNumber"]

	if err := c.Service.PatchEmployee(identityNumber, &employee); err != nil {
		log.Println("Error Updating employee:", err)
		http.Error(w, "Error updating employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee updated successfully"})
}
