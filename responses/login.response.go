package responses

import "personal_hr/models"

// LoginResponse ...
type LoginResponse struct {
	Token string `json:"token"`
}

// NewLoginResponse ...
func NewLoginResponse(person *models.Person) *LoginResponse {
	return &LoginResponse{
		Token: person.Token,
	}
}
