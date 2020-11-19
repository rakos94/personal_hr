package requests

import (
	"personal_hr/models"

	"github.com/labstack/echo"
)

// AddressRequest ...
type AddressRequest struct {
	PersonID   string  `form:"person_id" json:"person_id" validate:"required"`
	CityID     string  `form:"city_id" json:"city_id" validate:"required"`
	Address    string  `form:"address" json:"address" validate:"required"`
	PostalCode *string `form:"postal_code" json:"postal_code"`
	Phone      *string `form:"phone" json:"phone"`
	Fax        *string `form:"fax" json:"fax"`
	IsDefault  *bool   `form:"is_default" json:"is_default"`
	CreatedBy  string  `form:"created_by" json:"created_by" validate:"required"`
	UpdatedBy  string  `form:"updated_by" json:"updated_by" validate:"required"`
}

// Convert from echo FormValue
func (r AddressRequest) Convert(c echo.Context) *AddressRequest {
	c.Bind(&r)
	return &r
}

// Model from request
func (r AddressRequest) Model() *models.Address {
	return &models.Address{
		PersonID:   r.PersonID,
		CityID:     r.CityID,
		Address:    r.Address,
		PostalCode: r.PostalCode,
		Phone:      r.Phone,
		Fax:        r.Fax,
		IsDefault:  r.IsDefault,
		CreatedBy:  r.CreatedBy,
		UpdatedBy:  r.UpdatedBy,
	}
}
