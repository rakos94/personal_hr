package requests

import (
	"log"
	"personal_hr/models"

	"github.com/labstack/echo"
)

// EducationUpdateRequest ...
type EducationUpdateRequest struct {
	Institution string   `form:"institution" json:"institution"`
	Subject     string   `form:"subject" json:"subject"`
	Grade       *float64 `form:"grade" json:"grade"`
	YearBegin   *int     `form:"year_begin" json:"year_begin"`
	YearEnd     *int     `form:"year_end" json:"year_end"`
	UpdatedBy   string   `form:"updated_by" json:"updated_by"`
}

// Convert from echo FormValue
func (r EducationUpdateRequest) Convert(c echo.Context) *EducationUpdateRequest {
	if err := c.Bind(&r); err != nil {
		log.Println(err)
	}
	return &r
}

// Model from request
func (r EducationUpdateRequest) Model() *models.Education {
	return &models.Education{
		Institution: r.Institution,
		Subject:     r.Subject,
		Grade:       r.Grade,
		YearBegin:   r.YearBegin,
		YearEnd:     r.YearEnd,
		UpdatedBy:   r.UpdatedBy,
	}
}
