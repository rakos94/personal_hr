package requests

import (
	"personal_hr/helper"
	"personal_hr/models"

	"github.com/labstack/echo"
)

// PersonUpdateRequest ...
type PersonUpdateRequest struct {
	FirstName         string      `form:"first_name" json:"first_name"`
	LastName          string      `form:"last_name" json:"last_name"`
	BirthPlace        string      `form:"birth_place" json:"birth_place"`
	BirthDate         helper.Date `form:"birth_date" json:"birth_date"`
	Mobile            string      `form:"mobile" json:"mobile"`
	Gender            string      `form:"gender" json:"gender"`
	IsFromRecruitment *bool       `form:"is_from_recruitment" json:"is_from_recruitment"`
	UpdatedBy         string      `form:"updated_by" json:"updated_by"`
	// Email             string      `form:"email" validate:"email"`
}

// Convert from echo FormValue
func (p PersonUpdateRequest) Convert(c echo.Context) *PersonUpdateRequest {
	c.Bind(&p)
	return &p
}

// Model from request
func (p PersonUpdateRequest) Model() *models.Person {
	return &models.Person{
		FirstName:         p.FirstName,
		LastName:          &p.LastName,
		BirthPlace:        &p.BirthPlace,
		BirthDate:         p.BirthDate,
		Mobile:            &p.Mobile,
		Gender:            p.Gender,
		IsFromRecruitment: p.IsFromRecruitment,
		UpdatedBy:         p.UpdatedBy,
		// Email:             p.Email,
	}
}
