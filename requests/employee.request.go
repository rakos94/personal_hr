package requests

// EmployeeRequest ...
type EmployeeRequest struct {
	PersonID string `form:"person_id" validate:"required"`
}
