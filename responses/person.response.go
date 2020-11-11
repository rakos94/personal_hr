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
	rs.LastName = m.LastName.String
	rs.Email = m.Email
	rs.Gender = m.Gender.String
	rs.BaseCUModels = m.BaseCUModels
	return rs
}
