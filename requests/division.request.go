package requests

// DivisionRequest ...
type DivisionRequest struct {
	DivisionName string `form:"division_name" validate:"required"`
	DepartmentId string `form:"department_id" validate:"required"`
}
