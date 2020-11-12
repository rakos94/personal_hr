package responses

import (
	"personal_hr/models"
)

// PersonResponse ...
type PersonResponse struct {
	models.BaseModels
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     string  `json:"email"`
	Gender    *string `json:"gender"`
	models.BaseCUModels
}

// NewPersonResponse ...
func NewPersonResponse(m *models.Person) *PersonResponse {
	return &PersonResponse{
		BaseModels:   m.BaseModels,
		FirstName:    m.FirstName,
		LastName:     m.LastName,
		Email:        m.Email,
		Gender:       m.Gender,
		BaseCUModels: m.BaseCUModels,
	}
}

// NewListPersonResponse ...
func NewListPersonResponse(list []models.Person) []PersonResponse {
	var responses []PersonResponse
	for _, d := range list {
		responses = append(responses, *NewPersonResponse(&d))
	}
	return responses
}
