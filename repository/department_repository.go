package repository

import (
	"database/sql"
	"fmt"

	"github.com/gasBlar/GoGoManager/models"
)

type DepartmentRepository struct {
	DB *sql.DB
}

func NewDepartmentRepository(db *sql.DB) *DepartmentRepository {
	return &DepartmentRepository{DB: db}
}

func (r *DepartmentRepository) CreateDepartment(department *models.Department, profileId int) error {
	// Validasi apakah profileId ada di tabel profileManager
	var exists bool
	err := r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM profileManager WHERE id = ?)", profileId).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to validate profileId: %w", err)
	}
	if !exists {
		return fmt.Errorf("profileId %d does not exist", profileId)
	}

	// Insert department jika valid
	query := `INSERT INTO department (name, profileId) VALUES (?, ?)`
	result, err := r.DB.Exec(query, department.Name, profileId)
	if err != nil {
		return fmt.Errorf("failed to create department: %w", err)
	}

	// Ambil ID department yang baru dibuat
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}
	department.Id = int(lastInsertId)
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
	if department.ProfileId != nil {
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

func (r *DepartmentRepository) DeleteDepartment(departmenId string) error {

	query := `DELETE FROM department WHERE Id = ?`
	_, err := r.DB.Exec(query, departmenId)
	if err != nil {
		return err // Return the exact error for logging
	}
	return nil
}
