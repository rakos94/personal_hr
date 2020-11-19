package requests

import (
	"personal_hr/helper"
	"personal_hr/models"

	"github.com/labstack/echo"
)

// PersonRequest ...
type PersonRequest struct {
	FirstName         string      `form:"first_name" json:"first_name" validate:"required"`
	LastName          string      `form:"last_name" json:"last_name"`
	Email             string      `form:"email" json:"email" validate:"required,email"`
	Password          string      `form:"password" json:"password" validate:"required"`
	BirthPlace        string      `form:"birth_place" json:"birth_place"`
	BirthDate         helper.Date `form:"birth_date" json:"birth_date" validate:"required"`
	Mobile            string      `form:"mobile" json:"mobile"`
	Gender            string      `form:"gender" json:"gender" validate:"required"`
	IsFromRecruitment *bool       `form:"is_from_recruitment" json:"is_from_recruitment"`
	CreatedBy         string      `form:"created_by" json:"created_by" validate:"required"`
	UpdatedBy         string      `form:"updated_by" json:"updated_by" validate:"required"`
}

// Convert from echo FormValue
func (p PersonRequest) Convert(c echo.Context) *PersonRequest {
	c.Bind(&p)
	return &p
}

// Model from request
func (p PersonRequest) Model() *models.Person {
	return &models.Person{
		FirstName:         p.FirstName,
		LastName:          &p.LastName,
		Email:             p.Email,
		Password:          p.Password,
		BirthPlace:        &p.BirthPlace,
		BirthDate:         p.BirthDate,
		Mobile:            &p.Mobile,
		Gender:            p.Gender,
		IsFromRecruitment: p.IsFromRecruitment,
		CreatedBy:         p.CreatedBy,
		UpdatedBy:         p.UpdatedBy,
	}
}
