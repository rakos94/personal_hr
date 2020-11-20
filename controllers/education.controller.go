package controllers

import (
	"personal_hr/requests"
	"personal_hr/responses"
	"personal_hr/services"

	"github.com/labstack/echo"
)

var educationService services.EducationService = services.EducationServiceImpl{}

// SetEducation ...
func SetEducation(c *echo.Group) {
	c.POST("/education", createEducation)
	c.GET("/education", getEducationAll)
	c.GET("/education/:id", getEducationByID)
	c.PUT("/education/:id", updateEducation)
	c.GET("/person/:id/education", getEducationByPersonID)
}

func createEducation(c echo.Context) error {
	req := &requests.EducationRequest{}
	req = req.Convert(c)

	if err := c.Validate(req); err != nil {
		return resValErr(c, err)
	}

	result, err := educationService.CreateEducation(req.Model())
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewEducationResponse(result)
	return res(c, rs)
}

func getEducationAll(c echo.Context) error {
	result, err := educationService.GetEducationAll()
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewListEducationResponse(result)
	return res(c, rs)
}

func getEducationByID(c echo.Context) error {
	id := c.Param("id")

	result, err := educationService.GetEducationByID(id)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewEducationResponse(&result)
	return res(c, rs)
}

func updateEducation(c echo.Context) error {
	id := c.Param("id")
	req := &requests.EducationUpdateRequest{}
	req = req.Convert(c)

	if err := c.Validate(req); err != nil {
		return resValErr(c, err)
	}

	result, err := educationService.UpdateEducation(id, req.Model())
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewEducationResponse(result)
	return res(c, rs)
}

func getEducationByPersonID(c echo.Context) error {
	id := c.Param("id")

	result, err := educationService.GetEducationByPersonID(id)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewListEducationResponse(result)
	return res(c, rs)
}
