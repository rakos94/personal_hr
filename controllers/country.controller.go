package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"personal_hr/models"
	"personal_hr/services"
)

var countryService services.CountryService =services.CountryServiceImpl{}

func SetCountry(c *echo.Group)  {
	c.POST("/country", CreateCountry)
}
func CreateCountry(c echo.Context)error  {
	data := new(models.Country)
	if err:=  c.Bind(data);err !=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	err := countryService.CreateCountry(data)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.String(http.StatusOK,"ok")
}