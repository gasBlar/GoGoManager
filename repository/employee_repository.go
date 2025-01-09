package repository

import (
	"database/sql"
	"log"

	"github.com/gasBlar/GoGoManager/models"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) GetAllEmployees() ([]models.Employee, error) {
	// Query untuk mengambil semua data employee
	rows, err := r.DB.Query("SELECT identityNumber, name, employeeImageUri, gender, departmentId FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee

	// Iterasi hasil query dan masukkan ke dalam slice employees
	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(&employee.IdentityNumber, &employee.Name, &employee.EmployeeImageUri, &employee.Gender, &employee.DepartmentId); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
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

func (r *EmployeeRepository) DeleteEmployee(identityNumber string) error {
	query := `DELETE FROM employee WHERE identityNumber = ?`
	_, err := r.DB.Exec(query, identityNumber)
	if err != nil {
		return err // Return the exact error for logging
	}
	return nil
}

func (r *EmployeeRepository) PatchEmployee(identityNumber string, employee *models.EmployeePatch) error {
	// Buat query SQL secara dinamis berdasarkan kolom yang diubah
	query := "UPDATE employee SET "
	var args []interface{}
	// var updates []string

	if employee.IdentityNumber != nil {
		query += " identityNumber = ?,"
		args = append(args, *employee.IdentityNumber)
	}
	if employee.Name != nil {
		query += " name = ?,"
		args = append(args, *employee.Name)
	}
	if employee.EmployeeImageUri != nil {
		query += " employeeImageUri = ?,"
		args = append(args, *employee.EmployeeImageUri)
	}
	if employee.Gender != nil {
		query += " gender = ?,"
		args = append(args, *employee.Gender)
	}
	if employee.DepartmentId != nil {
		query += " departmentId = ?,"
		args = append(args, *employee.DepartmentId)
	}

	// Tambahkan kondisi untuk identityNumber
	query = query[:len(query)-1] + " WHERE identityNumber = ?"
	args = append(args, identityNumber)
	log.Println(query)

	_, err := r.DB.Exec(query, args...)
	if err != nil {
		return err // Return the exact error for logging
	}
	return nil
}
