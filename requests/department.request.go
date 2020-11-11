package requests

// DepartmentRequest ...
type DepartmentRequest struct {
	DepartmentName string `form:"department_name" validate:"required"`
	CompanyId      string `form:"company_id" validate:"required"`
}
