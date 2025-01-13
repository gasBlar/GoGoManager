package models

type Employee struct {
	IdentityNumber   string `json:"identityNumber" validate:"required,min=5,max=33"`
	Name             string `json:"name" validate:"required,min=4,max=33"`
	EmployeeImageUri string `json:"employeeImageUri" validate:"required,imagefileuri"`
	Gender           string `json:"gender" validate:"required,oneof=male female"`
	DepartmentId     string `json:"departmentId" validate:"required"`
}

type EmployeePatch struct {
	IdentityNumber   *string `json:"identityNumber" validate:"min=5,max=33"`
	Name             *string `json:"name" validate:"min=4,max=33"`
	EmployeeImageUri *string `json:"employeeImageUri" validate:"imagefileuri"`
	Gender           *string `json:"gender" validate:"oneof=male female"`
	DepartmentId     *string `json:"departmentId" validate:"gt=0"`
}
