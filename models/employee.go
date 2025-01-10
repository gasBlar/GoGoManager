package models

// type Employee struct {
// 	Id               int    `json:"id"`
// 	IdentityNumber   string `json:"identityNumber"`
// 	Name             string `json:"name"`
// 	EmployeeImageUri string `json:"employeeImageUri"`
// 	Gender           string `json:"gender"`
// 	DepartmentId     int    `json:"departmentId"`
// }

type Employee struct {
	Id               int    `json:"id"`
	IdentityNumber   string `json:"identityNumber" validate:"required,min=5,max=33"`
	Name             string `json:"name" validate:"required,min=4,max=33"`
	EmployeeImageUri string `json:"employeeImageUri" validate:"required,url"`
	Gender           string `json:"gender" validate:"required,oneof=Male Female"`
	DepartmentId     int    `json:"departmentId" validate:"required"`
}

// type EmployeePatch struct {
// 	IdentityNumber   *string `json:"identityNumber,omitempty"`
// 	Name             *string `json:"name,omitempty"`
// 	EmployeeImageUri *string `json:"employeeImageUri,omitempty"`
// 	Gender           *string `json:"gender,omitempty"`
// 	DepartmentId     *int    `json:"departmentId,omitempty"`
// }

type EmployeePatch struct {
	IdentityNumber   *string `json:"identityNumber" validate:"omitempty,min=5,max=33"`
	Name             *string `json:"name" validate:"omitempty,min=4,max=33"`
	EmployeeImageUri *string `json:"employeeImageUri" validate:"omitempty,url"`
	Gender           *string `json:"gender" validate:"omitempty,oneof=male female"`
	DepartmentId     *int    `json:"departmentId" validate:"omitempty,gt=0"`
}
