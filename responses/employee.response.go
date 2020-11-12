package responses

import "personal_hr/models"

// EmployeeResponse ...
type EmployeeResponse struct {
	models.BaseModels
	Nip      string          `json:"nip"`
	PersonID string          `json:"person_id"`
	Person   *PersonResponse ``
	models.BaseCUModels
}

// NewEmployeeResponse ...
func NewEmployeeResponse(m *models.Employee) *EmployeeResponse {
	return &EmployeeResponse{
		BaseModels:   m.BaseModels,
		Nip:          m.Nip,
		PersonID:     m.PersonID,
		Person:       NewPersonResponse(&m.Person),
		BaseCUModels: m.BaseCUModels,
	}
}
