package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"unicode/utf8"

	"net/http"

	"github.com/gasBlar/GoGoManager/api/v1/services"
	"github.com/gasBlar/GoGoManager/db"
	"github.com/gasBlar/GoGoManager/models"
	"github.com/gasBlar/GoGoManager/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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

	// **Validasi Content-Type**
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid Content-Type, must be application/json", http.StatusBadRequest)
		return
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate department name length
	nameLength := utf8.RuneCountInString(department.Name)
	if nameLength < 4 || nameLength > 33 {
		http.Error(w, "Department name must be between 4 and 33 characters", http.StatusBadRequest)
		return
	}

	// Call service to create department
	if err := c.Service.CreateDepartment(&department, user.Id); err != nil {
		log.Println("Error creating department:", err)
		http.Error(w, "Error creating department", http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{
		"departmentId": strconv.Itoa(department.Id),
		"name":         department.Name,
	}
	json.NewEncoder(w).Encode(response)
}

func (c *DepartmentController) PatchDepartment(w http.ResponseWriter, r *http.Request) {
	// user := r.Context().Value("user").(*utils.Claims)
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid Content-Type, must be application/json", http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	departmentId := vars["departmentId"]

	var department models.DepartmentPatch
	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		http.Error(w, "Invalid input b", http.StatusBadRequest)
		return
	}

	if len(department.Name) < 4 {
		utils.Response(w, http.StatusBadRequest, "", nil)
		return
	}

	err := validateUpdateRequestDep(department)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, fmt.Sprintf("Validation failed: %v", err), nil)
		return
	}

	database := db.DB
	var id string
	res := database.QueryRow("SELECT id FROM department WHERE id = ?", departmentId).Scan(&id)
	if res == sql.ErrNoRows {
		utils.Response(w, http.StatusNotFound, "Id", nil)
	}

	if err := c.Service.PatchDepartment(departmentId, &department); err != nil {
		log.Println("Error Updating department:", err)
		http.Error(w, "Error updating department", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"departmentId": departmentId,
		"name":         department.Name,
	}
	utils.Response(w, http.StatusOK, "", response)
}

func (c *DepartmentController) DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	departmentId := vars["departmentId"]
	// user := r.Context().Value("user").(*utils.Claims)

	var id string
	var name string
	res := db.DB.QueryRow("SELECT id, name FROM department WHERE id = ?", departmentId).Scan(&id, &name)
	if res == sql.ErrNoRows {
		utils.Response(w, http.StatusNotFound, "Id", nil)
		return
	}

	if err := c.Service.DeleteDepartment(departmentId); err != nil {
		if err.Error() == "access denied: manager does not have permission to modify this department" {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		log.Println("Error Deleting departmentawdasdaw:", err)
		http.Error(w, "Error deleting department", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"departmentId": id,
		"name":         name,
	}
	utils.Response(w, http.StatusOK, "", response)
}

func GetDepartments(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		departments, err := services.GetAllDepartments(r.Context(), db)
		if err != nil {
			http.Error(w, "Failed to retrieve departments", http.StatusInternalServerError)
			return
		}

		// Map the departments to the new structure
		type DepartmentResponse struct {
			DepartmentID int    `json:"departmentId"`
			Name         string `json:"name"`
		}
		var response []DepartmentResponse
		for _, dept := range departments {
			response = append(response, DepartmentResponse{
				DepartmentID: dept.Id,
				Name:         dept.Name,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func validateUpdateRequestDep(req models.DepartmentPatch) error {
	validate := validator.New()
	validate.RegisterValidation("nonempty", validateNonEmpty)

	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}
