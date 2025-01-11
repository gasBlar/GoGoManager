package services

import (
	"context"
	"database/sql"

	"github.com/gasBlar/GoGoManager/models"
	"github.com/gasBlar/GoGoManager/repository"
)

type DepartmentService struct {
	Repo *repository.DepartmentRepository
}

func NewDepartmentService(repo *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{Repo: repo}
}

func (s *DepartmentService) CreateDepartment(department *models.Department, id int) error {
	return s.Repo.CreateDepartment(department, id)
}

func GetAllDepartments(ctx context.Context, db *sql.DB) ([]models.Department, error) {
	query := "SELECT * FROM department"
	rows, err := db.QueryContext(ctx, query)
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
