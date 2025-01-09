package repository

import (
	"database/sql"

	"github.com/gasBlar/GoGoManager/models"
)

type DepartmentRepository struct {
	DB *sql.DB
}

func NewDepartmentRepository(db *sql.DB) *DepartmentRepository {
	return &DepartmentRepository{DB: db}
}

func (r *DepartmentRepository) CreateDepartment(department *models.Department, id int) error {
	query := `INSERT INTO department (name, ProfileId) 
              VALUES (?, ?)`
	_, err := r.DB.Exec(query, department.Name, id)
	if err != nil {
		return err // Return the exact error for logging
	}
	return nil
}