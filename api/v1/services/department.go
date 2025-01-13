package services

import (
	"context"
	"database/sql"
	"fmt"
	"unicode/utf8"

	"github.com/gasBlar/GoGoManager/models"
	"github.com/gasBlar/GoGoManager/repository"
)

type DepartmentService struct {
	Repo *repository.DepartmentRepository
}

func NewDepartmentService(repo *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{Repo: repo}
}

func (s *DepartmentService) CreateDepartment(department *models.Department, profileId int) error {
	// Validasi nama department
	nameLength := utf8.RuneCountInString(department.Name)
	if nameLength < 4 || nameLength > 33 {
		return fmt.Errorf("department name must be between 4 and 33 characters")
	}

	// Panggil repository untuk membuat department
	return s.Repo.CreateDepartment(department, profileId)
}

func GetAllDepartments(limit, offset int, name string, ctx context.Context, db *sql.DB) ([]models.Department, error) {
	query := "SELECT id, name, profileId FROM department WHERE 1=1"
	args := []interface{}{}

	if name != "" {
		query += " AND LOWER(name) LIKE ?"
		args = append(args, "%"+name+"%")
	}

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)
	// query := "SELECT * FROM department"
	rows, err := db.Query(query, args...)
	// rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []models.Department
	for rows.Next() {
		var department models.Department
		if err := rows.Scan(&department.Id, &department.Name, &department.ProfileId); err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}

func (s *DepartmentService) PatchDepartment(departmentId string, departmentPatch *models.DepartmentPatch) error {
	return s.Repo.PatchDepartment(departmentId, departmentPatch)
}

func (s *DepartmentService) DeleteDepartment(departmentId string) error {
	return s.Repo.DeleteDepartment(departmentId)
}
