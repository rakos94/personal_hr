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

var departmentService services.DepartmentService = services.DepartmentServiceImpl{}

// SetDepartment ...
func SetDepartment(c *echo.Group) {
	c.POST("/department", createDepartment)
	c.PUT("/department/:id", updateDepartment)
	c.GET("/department/:id", getDepartmentByID)
	c.GET("/department", getDepartmentAll)
	c.GET("/company/:id/department", getDepartmentByCompanyID)
}

func createDepartment(c echo.Context) error {
	req := &requests.DepartmentRequest{}

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return resValErr(c, err)
	}

	data := &models.Department{}
	helper.ConvertRequest(req, data)

	result, err := departmentService.CreateDepartment(data)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewDepartmentResponse(result)
	return res(c, rs)
}

func updateDepartment(c echo.Context) error {
	id := c.Param("id")
	req := &requests.DepartmentUpdateRequest{}

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return resValErr(c, err)
	}

	data := &models.Department{}
	helper.ConvertRequest(req, data)

	updated, err := departmentService.UpdateDepartment(id, data)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewDepartmentResponse(updated)
	return res(c, rs)
}

func getDepartmentByID(c echo.Context) error {
	id := c.Param("id")

	result, err := departmentService.GetDepartmentByID(id)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewDepartmentResponse(&result)
	return res(c, rs)
}

func getDepartmentAll(c echo.Context) error {
	result, err := departmentService.GetDepartmentAll()
	if err != nil {
		return resErr(c, err)
	}

	var resList []*responses.DepartmentResponse
	for _, data := range result {
		rs := responses.NewDepartmentResponse(&data)
		resList = append(resList, rs)
	}

	return res(c, resList)
}

func getDepartmentByCompanyID(c echo.Context) error {
	id := c.Param("id")
	result, err := departmentService.GetDepartmentByCompanyID(id)
	if err != nil {
		return resErr(c, err)
	}

	var resList []*responses.DepartmentResponse
	for _, data := range result {
		rs := responses.NewDepartmentResponse(&data)
		resList = append(resList, rs)
	}

	return res(c, resList)
}
