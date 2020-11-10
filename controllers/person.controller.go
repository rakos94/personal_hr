package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"personal_hr/configs"
	"personal_hr/models"
	"personal_hr/requests"
	"personal_hr/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var personService services.PersonService = services.PersonServiceImpl{}

func SetPerson(c *echo.Group, e *echo.Echo) {
	e.POST("/person", createPerson)
	e.POST("/api/login", login)
	c.GET("/person/:id", getPersonById)
	c.GET("/user", getUser)
}

func createPerson(c echo.Context) error {
	data := &models.Person{}

	if err := c.Bind(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(data); err != nil {
		return resValErr(c, err)
	}

	result, err := personService.CreatePerson(data)
	if err == nil {
		return res(c, result)
	}

	return resErr(c, err)
}

func login(c echo.Context) error {
	data := &requests.Login{}

	if err := c.Bind(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(data); err != nil {
		return resValErr(c, err)
	}

	var result, err = personService.Login(data.Email, data.Password)
	if err == nil {
		v, _ := configs.CreateJwtToken(data.Email)
		result.Token = v
		return res(c, result)
	}

	return resErr(c, err)
}

func getPersonById(c echo.Context) error {
	id := c.Param("id")

	var result, err = personService.GetPersonById(id)
	if err == nil {
		if result.Count != 0 {
			return res(c, result)
		}

		return resErr(c, errors.New("Data not found"))
	}

	return resErr(c, err)
}

func getUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	return res(c, claims)
}
