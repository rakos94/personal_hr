package requests

// DepartmentRequest ...
type DepartmentRequest struct {
	DepartmentName string `form:"department_name" validate:"required"`
	CompanyID      string `form:"company_id" validate:"required"`
}

// DepartmentUpdateRequest ...
type DepartmentUpdateRequest struct {
	DepartmentName string `form:"department_name" validate:"required"`
}
