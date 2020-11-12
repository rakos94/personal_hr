package controllers

import (
	"fmt"
	"log"
	"net/http"
	"personal_hr/helper"
	"personal_hr/models"
	"personal_hr/requests"
	"personal_hr/responses"
	"personal_hr/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var personService services.PersonService = services.PersonServiceImpl{}

// SetPerson ...
func SetPerson(c *echo.Group, e *echo.Echo) {
	e.POST("/api/login", login)
	e.POST("/person", createPerson)
	c.GET("/person", getPersonAll)
	c.GET("/person/:id", getPersonByID)
	c.PUT("/person/:id", updatePerson)
	c.PUT("/person/:id/password-update", updatePassword)
	c.GET("/user", getUser)
}

func login(c echo.Context) error {
	data := &requests.Login{}

	if err := c.Bind(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(data); err != nil {
		return resValErr(c, err)
	}

	result, err := personService.Login(data.Email, data.Password)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewLoginResponse(&result)
	return res(c, rs)
}

func createPerson(c echo.Context) error {
	req := &requests.PersonRequest{}

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return resValErr(c, err)
	}

	data := &models.Person{}
	helper.ConvertRequest(req, data)

	result, err := personService.CreatePerson(data)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewPersonResponse(result)
	return res(c, rs)
}

func getPersonAll(c echo.Context) error {
	result, err := personService.GetPersonAll()
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewListPersonResponse(result)
	return res(c, rs)
}

func getPersonByID(c echo.Context) error {
	id := c.Param("id")

	result, err := personService.GetPersonByID(id)
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewPersonResponse(&result)
	return res(c, rs)
}

func updatePerson(c echo.Context) error {
	// id := c.Param("id")
	// req := &requests.PersonUpdateRequest{}

	log.Println(c.FormValue("birth_date"))
	// if err := c.Bind(req); err != nil {
	// 	log.Println("Bind")
	// 	log.Println(err)
	// 	return c.String(http.StatusBadRequest, err.Error())
	// }

	// if err := c.Validate(req); err != nil {
	// 	return resValErr(c, err)
	// }

	// data := &models.Person{}
	// data.ConvertFromRequest(req)
	// // req.NewPersonModel()
	// log.Println(req.BirthDate)
	// // helper.ConvertRequest(req, data)
	// log.Println(data.BirthDate)

	// updated, err := personService.UpdatePerson(id, data)
	// if err != nil {
	// 	return resErr(c, err)
	// }

	// rs := responses.NewPersonResponse(updated)
	rs := "Test"
	return res(c, rs)
}

func updatePassword(c echo.Context) error {
	id := c.Param("id")
	req := &requests.PersonPasswordRequest{}

	if err := c.Bind(req); err != nil {
		return resErr(c, err)
	}

	if err := c.Validate(req); err != nil {
		return resValErr(c, err)
	}

	err := personService.UpdatePassword(id, req.Password)
	if err != nil {
		return resErr(c, err)
	}

	return res(c, "Success update password")
}

func getUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	return res(c, claims)
}
