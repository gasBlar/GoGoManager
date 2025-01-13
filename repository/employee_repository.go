package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/gasBlar/GoGoManager/models"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

// Fungsi untuk memvalidasi apakah manager memiliki akses ke employee
func (r *EmployeeRepository) ValidateManagerAccess(managerId int, identityNumber string) error {
	query := `
        SELECT COUNT(*)
        FROM employee e
        JOIN department d ON e.departmentId = d.id
        JOIN profileManager pm ON d.profileId = pm.id
        WHERE e.identityNumber = ? AND pm.id = ?`

	var count int
	err := r.DB.QueryRow(query, identityNumber, managerId).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("access denied: manager does not have permission to modify this employee")
	}

	return nil
}

func (r *EmployeeRepository) GetAllEmployees(limit, offset int, identityNumber, name, gender, departmentId string) ([]models.Employee, error) {
	// Query untuk mengambil semua data employee
	// query := `SELECT identityNumber, name, employeeImageUri, gender, departmentId
	// 		FROM employee`

	query := "SELECT identityNumber, name, employeeImageUri, gender, departmentId FROM employee WHERE 1=1"
	args := []interface{}{}

	if identityNumber != "" {
		query += " AND LOWER(identityNumber) LIKE ?"
		args = append(args, identityNumber+"%")
	}

	if name != "" {
		query += " AND LOWER(name) LIKE ?"
		args = append(args, "%"+name+"%")
	}

	if gender != "" {
		query += " AND gender = ?"
		args = append(args, gender)
	}

	if departmentId != "" {
		query += " AND departmentId = ?"
		args = append(args, departmentId)
	}

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
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

func (r *EmployeeRepository) CreateEmployee(employee *models.Employee) ([]models.Employee, error) {
	query := `INSERT INTO employee (identityNumber, name, employeeImageUri, gender, departmentId) 
              VALUES (?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, employee.IdentityNumber, employee.Name, employee.EmployeeImageUri, employee.Gender, employee.DepartmentId)
	if err != nil {
		return nil, err // Return the exact error for logging
	}
	return nil, err
}

func (r *EmployeeRepository) DeleteEmployee(identityNumber string) error {
	queryValidate := `SELECT COUNT(*) FROM employee WHERE identityNumber = ?`
	var count int
	errVal := r.DB.QueryRow(queryValidate, identityNumber).Scan(&count)
	if errVal != nil {
		return errVal
	}
	if count == 0 {
		return errors.New("identityNumber not found")
	}

	query := `DELETE FROM employee WHERE identityNumber = ?`
	_, err := r.DB.Exec(query, identityNumber)
	if err != nil {
		return err // Return the exact error for logging
	}
	return nil
}

func (r *EmployeeRepository) PatchEmployee(identityNumber string, employee *models.EmployeePatch) (string, error) {
	queryValidate := `SELECT COUNT(*) FROM employee WHERE identityNumber = ?`
	var count int
	errVal := r.DB.QueryRow(queryValidate, identityNumber).Scan(&count)
	if errVal != nil {
		return "", errVal
	}
	if count == 0 {
		return "", errors.New("identityNumber not found")
	}

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

	_, err := r.DB.Exec(query, args...)
	if err != nil {
		return "", err // Return the exact error for logging
	}

	log.Println()

	// Jika identityNumber diubah, gunakan identityNumber baru, jika tidak, gunakan yang lama
	if employee.IdentityNumber != nil {
		return *employee.IdentityNumber, nil
	}
	return identityNumber, nil
}

func (r *EmployeeRepository) GetEmployeeByIdentityNumber(identityNumber string) (*models.Employee, error) {
	// // Validasi akses manager
	// if err := r.ValidateManagerAccess(managerId, identityNumber); err != nil {
	// 	return nil, err
	// }

	// Ambil data employee berdasarkan identityNumber
	query := `
		SELECT identityNumber, name, employeeImageUri, gender, departmentId
		FROM employee
		WHERE identityNumber = ?
	`
	var employee models.Employee
	err := r.DB.QueryRow(query, identityNumber).Scan(
		&employee.IdentityNumber,
		&employee.Name,
		&employee.EmployeeImageUri,
		&employee.Gender,
		&employee.DepartmentId,
	)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}
