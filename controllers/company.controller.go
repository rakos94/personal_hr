package controllers

import (
	"net/http"
	"personal_hr/helper"
	"personal_hr/models"
	"personal_hr/requests"
	"personal_hr/responses"
	"personal_hr/services"

	"github.com/labstack/echo"
)

var companyService services.CompanyService = services.CompanyServiceImpl{}

// SetCompany ...
func SetCompany(c *echo.Group) {
	c.POST("/company", createCompany)
	c.PUT("/company/:id", updateCompany)
	c.GET("/company/:id", getCompanyByID)
	c.GET("/company", getCompanyAll)
}

func createCompany(c echo.Context) error {
	req := &requests.CompanyRequest{}

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return resValErr(c, err)
	}

	data := &models.Company{}
	helper.ConvertRequest(req, data)

	result, err := companyService.CreateCompany(data)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewCompanyResponse(result)
	return res(c, rs)
}

func updateCompany(c echo.Context) error {
	id := c.Param("id")
	req := &requests.CompanyRequest{}

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return resValErr(c, err)
	}

	data := &models.Company{}
	helper.ConvertRequest(req, data)

	updated, err := companyService.UpdateCompany(id, data)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewCompanyResponse(&updated)
	return res(c, rs)
}

func getCompanyByID(c echo.Context) error {
	id := c.Param("id")

	result, err := companyService.GetCompanyByID(id)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewCompanyResponse(&result)
	return res(c, rs)
}

func getCompanyAll(c echo.Context) error {
	result, err := companyService.GetCompanyAll()
	if err != nil {
		return resErr(c, err)
	}

	var resList []*responses.CompanyResponse
	for _, data := range result {
		rs := responses.NewCompanyResponse(&data)
		resList = append(resList, rs)
	}

	return res(c, resList)
}
