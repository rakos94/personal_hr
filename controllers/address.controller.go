package controllers

import (
	"personal_hr/requests"
	"personal_hr/responses"
	"personal_hr/services"

	"github.com/labstack/echo"
)

var addressService services.AddressService = services.AddressServiceImpl{}

// SetAddress ...
func SetAddress(c *echo.Group) {
	c.POST("/address", createAddress)
	c.GET("/address", getAddressAll)
	c.GET("/address/:id", getAddressByID)
	c.PUT("/address/:id", updateAddress)
	c.GET("/person/:id/address", getAddressByPersonID)
}

func createAddress(c echo.Context) error {
	req := &requests.AddressRequest{}
	req = req.Convert(c)

	if err := c.Validate(req); err != nil {
		return resValErr(c, err)
	}

	result, err := addressService.CreateAddress(req.Model())
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewAddressResponse(result)
	return res(c, rs)
}

func getAddressAll(c echo.Context) error {
	result, err := addressService.GetAddressAll()
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewListAddressResponse(result)
	return res(c, rs)
}

func getAddressByID(c echo.Context) error {
	id := c.Param("id")

	result, err := addressService.GetAddressByID(id)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewAddressResponse(&result)
	return res(c, rs)
}

func getAddressByPersonID(c echo.Context) error {
	id := c.Param("id")

	result, err := addressService.GetAddressByPersonID(id)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewListAddressResponse(result)
	return res(c, rs)
}

func updateAddress(c echo.Context) error {
	id := c.Param("id")
	req := &requests.AddressUpdateRequest{}
	req = req.Convert(c)

	if err := c.Validate(req); err != nil {
		return resValErr(c, err)
	}

	result, err := addressService.UpdateAddress(id, req.Model())
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewAddressResponse(result)
	return res(c, rs)
}
