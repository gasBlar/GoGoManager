package repository

import (
	"database/sql"

	"github.com/gasBlar/GoGoManager/models"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) CreateEmployee(employee *models.Employee) error {
	query := `INSERT INTO employee (identityNumber, name, employeeImageUri, gender, departmentId) 
              VALUES (?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, employee.IdentityNumber, employee.Name, employee.EmployeeImageUri, employee.Gender, employee.DepartmentId)
	if err != nil {
		return err // Return the exact error for logging
	}
	return nil
}
