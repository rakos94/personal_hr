package responses

import (
	"personal_hr/models"
)

// PersonResponse ...
type PersonResponse struct {
	models.BaseModels
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	models.BaseCUModels
}

// NewPersonResponse ...
func NewPersonResponse(m *models.Person) *PersonResponse {
	rs := &PersonResponse{}
	rs.BaseModels = m.BaseModels
	rs.FirstName = m.FirstName
	if m.LastName != nil {
		rs.LastName = *m.LastName
	}
	rs.Email = m.Email
	if m.Gender != nil {
		rs.Gender = *m.Gender
	}
	rs.BaseCUModels = m.BaseCUModels
	return rs
}
