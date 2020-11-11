package requests

// EmployeeRequest ...
type EmployeeRequest struct {
	PersonId string `form:"person_id" json:"person_id" validate:"required"`
}
