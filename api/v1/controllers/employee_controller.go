package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

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
	//validasi content type
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Unsupported Content-Type, must be application/json", http.StatusBadRequest)
		return
	}

	var employee models.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		log.Println(err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := utils.ValidateAddEmployee(employee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdEmployee, err := c.Service.CreateEmployee(&employee)
	if err != nil {
		if strings.Contains(fmt.Sprint(err), "Error 1062 (23000)") {
			http.Error(w, "Identity Number must be unique", http.StatusConflict)
			return
		}
		if strings.Contains(fmt.Sprint(err), "a foreign key constraint fails") {
			http.Error(w, "Department Id Not Valid", http.StatusBadRequest)
			return
		}
		log.Println(err)
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
	// Parse query parameters
	queryParams := r.URL.Query()

	// Default values
	limit := 5
	offset := 0

	// Parse limit and offset
	if l, err := strconv.Atoi(queryParams.Get("limit")); err == nil && l > 0 {
		limit = l
	}
	if o, err := strconv.Atoi(queryParams.Get("offset")); err == nil && o >= 0 {
		offset = o
	}

	// Parse other parameters
	identityNumber := queryParams.Get("identityNumber")
	name := queryParams.Get("name")
	gender := queryParams.Get("gender")
	departmentId := queryParams.Get("departmentId")

	// Mendapatkan data seluruh employee dari service
	employees, err := c.Service.GetAllEmployees(limit, offset, identityNumber, name, gender, departmentId)
	if err != nil {
		log.Println(err)
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
	// user := r.Context().Value("user").(*utils.Claims)

	if err := c.Service.DeleteEmployee(identityNumber); err != nil {
		if strings.Contains(fmt.Sprint(err), "identityNumber not found") {
			http.Error(w, "identityNumber not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error deleting employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted"})
}

func (c *EmployeeController) PatchEmployee(w http.ResponseWriter, r *http.Request) {
	// Validasi Content-Type header
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Unsupported Content-Type, must be application/json", http.StatusBadRequest)
		return
	}

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

	log.Println((employee.IdentityNumber))

	newIdentityNumber, err := c.Service.PatchEmployee(identityNumber, &employee)
	if err != nil {
		if strings.Contains(fmt.Sprint(err), "Error 1062 (23000)") {
			http.Error(w, "Identity Number is used", http.StatusConflict)
			return
		}
		if strings.Contains(fmt.Sprint(err), "a foreign key constraint fails") {
			http.Error(w, "Department Id Not Valid", http.StatusBadRequest)
			return
		}
		if strings.Contains(fmt.Sprint(err), "identityNumber not found") {
			http.Error(w, "identityNumber not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error updating employee", http.StatusInternalServerError)
		return
	}

	log.Println(newIdentityNumber)

	// Ambil data employee yang baru diperbarui
	updatedEmployee, err := c.Service.GetEmployeeByIdentityNumber(newIdentityNumber)
	if err != nil {
		http.Error(w, "Error fetching updated employee", http.StatusInternalServerError)
		return
	}

	// Kirimkan response JSON dengan format sesuai permintaan
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEmployee)

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(map[string]string{"message": "Employee updated successfully"})
}
