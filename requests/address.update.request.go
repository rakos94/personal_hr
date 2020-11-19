package requests

import (
	"personal_hr/models"

	"github.com/labstack/echo"
)

// AddressUpdateRequest ...
type AddressUpdateRequest struct {
	CityID     string  `form:"city_id" json:"city_id" validate:"required"`
	Address    string  `form:"address" json:"address" validate:"required"`
	PostalCode *string `form:"postal_code" json:"postal_code"`
	Phone      *string `form:"phone" json:"phone"`
	Fax        *string `form:"fax" json:"fax"`
	IsDefault  *bool   `form:"is_default" json:"is_default"`
	UpdatedBy  string  `form:"updated_by" json:"updated_by" validate:"required"`
}

// Convert from echo FormValue
func (r AddressUpdateRequest) Convert(c echo.Context) *AddressUpdateRequest {
	c.Bind(&r)
	return &r
}

// Model from request
func (r AddressUpdateRequest) Model() *models.Address {
	return &models.Address{
		CityID:     r.CityID,
		Address:    r.Address,
		PostalCode: r.PostalCode,
		Phone:      r.Phone,
		Fax:        r.Fax,
		IsDefault:  r.IsDefault,
		UpdatedBy:  r.UpdatedBy,
	}
}
