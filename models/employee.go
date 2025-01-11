package models

type Employee struct {
	IdentityNumber   string `json:"identityNumber" validate:"required,min=5,max=33"`
	Name             string `json:"name" validate:"required,min=4,max=33"`
	EmployeeImageUri string `json:"employeeImageUri" validate:"required,url"`
	Gender           string `json:"gender" validate:"required,oneof=male female"`
	DepartmentId     int    `json:"departmentId" validate:"required"`
}

type EmployeePatch struct {
	IdentityNumber   *string `json:"identityNumber" validate:"omitempty,min=5,max=33"`
	Name             *string `json:"name" validate:"omitempty,min=4,max=33"`
	EmployeeImageUri *string `json:"employeeImageUri" validate:"omitempty,url"`
	Gender           *string `json:"gender" validate:"omitempty,oneof=male female"`
	DepartmentId     *int    `json:"departmentId" validate:"omitempty,gt=0"`
}
