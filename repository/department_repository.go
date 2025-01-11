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
	query := `INSERT INTO department (name, profileId) 
              VALUES (?, ?)`
	_, err := r.DB.Exec(query, department.Name, id)
	if err != nil {
		return err // Return the exact error for logging
	}
	return nil
}

func (r *DepartmentRepository) PatchDepartment(departmentId string, department *models.DepartmentPatch) error {
	// Buat query SQL secara dinamis berdasarkan kolom yang diubah
	query := "UPDATE department SET "
	var args []interface{}

	if department.Name != "" {
		query += " name = ?,"
		args = append(args, department.Name)
	}
	if department.ProfileId != "" {
		query += " profileId = ?,"
		args = append(args, department.ProfileId)
	}

	// Tambahkan kondisi untuk departmentId
	query = query[:len(query)-1] + " WHERE Id = ?"
	args = append(args, departmentId)
	// log.Println(query)

	_, err := r.DB.Exec(query, args...)
	if err != nil {
		return err // Return the exact error for logging
	}
	return nil
}

func (r *DepartmentRepository) DeleteDepartment(managerId int, identityNumber string) error {
	// if err := r.ValidateManagerAccess(managerId, identityNumber); err != nil {
	// 	return err
	// }

	query := `DELETE FROM department WHERE Id = ?`
	_, err := r.DB.Exec(query, identityNumber)
	if err != nil {
		return err // Return the exact error for logging
	}
	return nil
}
