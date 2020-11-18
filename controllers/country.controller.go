package controllers

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"personal_hr/models"
	"personal_hr/services"
)

var countryService services.CountryService =services.CountryServiceImpl{}

func SetCountry(c *echo.Group)  {
	c.POST("/country", CreateCountry)
	c.GET("/country",GetCountryAll)
	c.GET("/country/:id",GetCountryById)
	c.PUT("/country/:id",UpdateCountry)
	c.DELETE("/country/:id",DeleteCountry)
	c.GET("/country/name/:name",GetCountryByName)

}
func CreateCountry(c echo.Context)error  {
	data := new(models.Country)
	if err:=  c.Bind(data);err !=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	err := countryService.CreateCountry(data)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func GetCountryAll(c echo.Context)(error)  {
	data,err := countryService.GetCountryAll()
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func GetCountryById(c echo.Context)(error)  {
	id:= c.Param("id")

	data,err:=countryService.GetCountryById(id)
	if err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func UpdateCountry(c echo.Context)error  {
	id:= c.Param("id")
	data := new(models.Country)
	if err:=  c.Bind(data);err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	err:=countryService.UpdateCountry(id,data)
	if err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func DeleteCountry(c echo.Context)error  {
	id:=c.Param("id")
	err:= countryService.DeleteCountry(id)
	if err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,"delete done")
}
func GetCountryByName(c echo.Context)error  {
	name:=c.Param("name")
	log.Println("name :",name)
	data,err := countryService.GetCountryByName(name)
	if err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}