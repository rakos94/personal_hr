package requests

// DivisionRequest ...
type DivisionRequest struct {
	DivisionName string `json:"division_name" validate:"required"`
	DepartmentId string `json:"department_id" validate:"required"`
}
