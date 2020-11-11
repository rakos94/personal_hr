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
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewEmployeeResponse(result)
	return res(c, rs)
}

func getEmployeeByID(c echo.Context) error {
	id := c.Param("id")

	result, err := employeeService.GetEmployeeByID(id)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewEmployeeResponse(&result)
	return res(c, rs)
}

func getEmployeeAll(c echo.Context) error {
	result, err := employeeService.GetEmployeeAll()
	if err != nil {
		return resErr(c, err)
	}

	var resList []*responses.EmployeeResponse
	for _, data := range result {
		rs := responses.NewEmployeeResponse(&data)
		resList = append(resList, rs)
	}
	return res(c, result)
}
