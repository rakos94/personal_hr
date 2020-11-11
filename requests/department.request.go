package requests

// DepartmentRequest ...
type DepartmentRequest struct {
	DepartmentName string `json:"department_name" validate:"required"`
	CompanyId      string `json:"company_id" validate:"required"`
}
