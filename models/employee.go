package models

type Employee struct {
	Id               int    `json:"id"`
	IdentityNumber   string `json:"identityNumber"`
	Name             string `json:"name"`
	EmployeeImageUri string `json:"employeeImageUri"`
	Gender           string `json:"gender"`
	DepartmentId     int    `json:"departmentId"`
}
