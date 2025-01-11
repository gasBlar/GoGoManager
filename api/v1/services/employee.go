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

func (s *EmployeeService) CreateEmployee(employee *models.Employee) (*models.Employee, error) {
	_, err := s.Repo.CreateEmployee(employee)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *EmployeeService) GetAllEmployees(managerId int) ([]models.Employee, error) {
	return s.Repo.GetAllEmployees(managerId)
}

func (s *EmployeeService) DeleteEmployee(userId int, identityNumber string) error {
	return s.Repo.DeleteEmployee(userId, identityNumber)
}

func (s *EmployeeService) PatchEmployee(managerId int, identityNumber string, employeePatch *models.EmployeePatch) (string, error) {
	return s.Repo.PatchEmployee(managerId, identityNumber, employeePatch)
}

func (s *EmployeeService) GetEmployeeByIdentityNumber(managerId int, identityNumber string) (*models.Employee, error) {
	return s.Repo.GetEmployeeByIdentityNumber(managerId, identityNumber)
}
