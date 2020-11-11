package controllers

import (
	"net/http"
	"personal_hr/helper"
	"personal_hr/models"
	"personal_hr/requests"
	"personal_hr/services"

	"github.com/labstack/echo"
)

var employeeService services.EmployeeService = services.EmployeeServiceImpl{}

// SetEmployee ...
func SetEmployee(c *echo.Group) {
	c.POST("/employee", createEmployee)
	c.GET("/employee/:id", getEmployeeByID)
	c.GET("/employee", getEmployeeAll)
}

func createEmployee(c echo.Context) error {
	req := &requests.EmployeeRequest{}

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return resValErr(c, err)
	}

	data := &models.Employee{}
	helper.ConvertRequest(req, data)

	result, err := employeeService.CreateEmployee(data)
	if err == nil {
		return res(c, result)
	}

	return resErr(c, err)
}

func getEmployeeByID(c echo.Context) error {
	id := c.Param("id")

	var result, err = employeeService.GetEmployeeByID(id)
	if err == nil {
		return res(c, result)
	}

	return resErr(c, err)
}

func getEmployeeAll(c echo.Context) error {
	var result, err = employeeService.GetEmployeeAll()
	if err == nil {
		return res(c, result)
	}

	return resErr(c, err)
}
