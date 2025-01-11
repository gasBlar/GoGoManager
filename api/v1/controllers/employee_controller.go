package controllers

import (
	"encoding/json"
	"log"

	"net/http"

	"github.com/gasBlar/GoGoManager/api/v1/services"
	"github.com/gasBlar/GoGoManager/models"
	"github.com/gasBlar/GoGoManager/utils"
	"github.com/gorilla/mux"
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

	if err := utils.ValidateAddEmployee(employee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdEmployee, err := c.Service.CreateEmployee(&employee)
	if err != nil {
		http.Error(w, "Error creating employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdEmployee)

	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(map[string]string{"message": "Employee created successfully"})
}

func (c *EmployeeController) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*utils.Claims)
	managerId := user.Id
	// Mendapatkan data seluruh employee dari service
	employees, err := c.Service.GetAllEmployees(managerId)
	if err != nil {
		http.Error(w, "Error retrieving employees", http.StatusInternalServerError)
		return
	}

	// Menyusun respons dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func (c *EmployeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	identityNumber := vars["identityNumber"]
	user := r.Context().Value("user").(*utils.Claims)

	if err := c.Service.DeleteEmployee(user.Id, identityNumber); err != nil {
		if err.Error() == "access denied: manager does not have permission to modify this employee" {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, "Error deleting employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee deleted successfully"})
}

func (c *EmployeeController) PatchEmployee(w http.ResponseWriter, r *http.Request) {
	// Ambil managerId dari token (contoh dengan header Authorization)
	user := r.Context().Value("user").(*utils.Claims)
	log.Println(user.Id)

	var employee models.EmployeePatch
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Ambil identityNumber dari URL
	vars := mux.Vars(r)
	identityNumber := vars["identityNumber"]

	if err := utils.ValidateAddEmployee(employee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Service.PatchEmployee(user.Id, identityNumber, &employee); err != nil {
		if err.Error() == "access denied: manager does not have permission to modify this employee" {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		log.Println("Error Updating employee:", err)
		http.Error(w, "Error updating employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee updated successfully"})
}
