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
func NewPersonResponse(person *models.Person) *PersonResponse {
	rs := &PersonResponse{}
	rs.BaseModels = person.BaseModels
	rs.FirstName = person.FirstName
	rs.LastName = person.LastName.String
	rs.Email = person.Email
	rs.Gender = person.Gender.String
	rs.BaseCUModels = person.BaseCUModels
	return rs
}
