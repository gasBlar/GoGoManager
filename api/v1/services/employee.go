package services

import (
	"github.com/gasBlar/GoGoManager/models"
	"github.com/gasBlar/GoGoManager/repository"
)

type EmployeeService struct {
	Repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{Repo: repo}
}

func (s *EmployeeService) CreateEmployee(employee *models.Employee) error {
	return s.Repo.CreateEmployee(employee)
}

func (s *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.Repo.GetAllEmployees()
}

// func GetAllEmployees(ctx context.Context, db *sql.DB) ([]models.Employee, error) {
// 	query := "SELECT * FROM employee"
// 	rows, err := db.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var employees []models.Employee
// 	for rows.Next() {
// 		var employee models.Employee
// 		if err := rows.Scan(&employee.Id, &employee.IdentityNumber, &employee.Name, &employee.EmployeeImageUri, &employee.Gender, &employee.DepartmentId); err != nil {
// 			return nil, err
// 		}
// 		employees = append(employees, employee)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return employees, nil
// }

func (s *EmployeeService) DeleteEmployee(userId int, identityNumber string) error {
	return s.Repo.DeleteEmployee(userId, identityNumber)
}

func (s *EmployeeService) PatchEmployee(managerId int, identityNumber string, employeePatch *models.EmployeePatch) error {
	return s.Repo.PatchEmployee(managerId, identityNumber, employeePatch)
}
