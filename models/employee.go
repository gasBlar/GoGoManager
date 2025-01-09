package models

type Employee struct {
	Id               int    `json:"id"`
	IdentityNumber   string `json:"identityNumber"`
	Name             string `json:"name"`
	EmployeeImageUri string `json:"employeeImageUri"`
	Gender           string `json:"gender"`
	DepartmentId     int    `json:"departmentId"`
}

type EmployeePatch struct {
	IdentityNumber   *string `json:"identityNumber,omitempty"`
	Name             *string `json:"name,omitempty"`
	EmployeeImageUri *string `json:"employeeImageUri,omitempty"`
	Gender           *string `json:"gender,omitempty"`
	DepartmentId     *int    `json:"departmentId,omitempty"`
}
