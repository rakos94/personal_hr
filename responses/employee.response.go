package responses

import "personal_hr/models"

// EmployeeResponse ...
type EmployeeResponse struct {
	models.BaseModels
	Nip      string        `json:"nip"`
	PersonID string        `json:"person_id"`
	Person   models.Person ``
	models.BaseCUModels
}

// NewEmployeeResponse ...
func NewEmployeeResponse(m *models.Employee) *EmployeeResponse {
	return &EmployeeResponse{
		BaseModels:   m.BaseModels,
		Nip:          m.Nip,
		PersonID:     m.PersonID,
		Person:       m.Person,
		BaseCUModels: m.BaseCUModels,
	}
}
