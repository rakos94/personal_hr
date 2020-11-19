package requests

import (
	"personal_hr/models"

	"github.com/labstack/echo"
)

// EducationRequest ...
type EducationRequest struct {
	PersonID    string   `form:"person_id" json:"person_id" validate:"required"`
	Institution string   `form:"institution" json:"institution" validate:"required"`
	Subject     string   `form:"subject" json:"subject" validate:"required"`
	Grade       *float64 `form:"grade" json:"grade" validate:"required"`
	YearBegin   *int     `form:"year_begin" json:"year_begin"`
	YearEnd     *int     `form:"year_end" json:"year_end"`
	CreatedBy   string   `form:"created_by" json:"created_by" validate:"required"`
	UpdatedBy   string   `form:"updated_by" json:"updated_by" validate:"required"`
}

// Convert from echo FormValue
func (r EducationRequest) Convert(c echo.Context) *EducationRequest {
	c.Bind(&r)
	return &r
}

// Model from request
func (r EducationRequest) Model() *models.Education {
	return &models.Education{
		PersonID:    r.PersonID,
		Institution: r.Institution,
		Subject:     r.Subject,
		Grade:       r.Grade,
		YearBegin:   r.YearBegin,
		YearEnd:     r.YearEnd,
		CreatedBy:   r.CreatedBy,
		UpdatedBy:   r.UpdatedBy,
	}
}
