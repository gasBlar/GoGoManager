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

func (s *EmployeeService) GetAllEmployees(managerId int) ([]models.Employee, error) {
	return s.Repo.GetAllEmployees(managerId)
}

func (s *EmployeeService) DeleteEmployee(userId int, identityNumber string) error {
	return s.Repo.DeleteEmployee(userId, identityNumber)
}

func (s *EmployeeService) PatchEmployee(managerId int, identityNumber string, employeePatch *models.EmployeePatch) error {
	return s.Repo.PatchEmployee(managerId, identityNumber, employeePatch)
}
