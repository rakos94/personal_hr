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
	c.GET("/country",GetCountryAll)
	c.GET("/country/:id",GetCountryById)
	c.PUT("/country/:id",PutCountry)
	c.DELETE("/country/:id",DeleteCountry)

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
	return c.JSON(http.StatusOK,err)
}
func GetCountryAll(c echo.Context)(error)  {
	data,err := countryService.GetCountryAll()
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusOK,data)
}
func GetCountryById(c echo.Context)(error)  {
	id:= c.Param("id")

	data,err:=countryService.GetCountryById(id)
	if err !=nil{
		return c.String(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func PutCountry(c echo.Context)error  {
	id:= c.Param("id")
	data := new(models.Country)
	if err:=  c.Bind(data);err !=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	err:=countryService.PutCountry(id,data)
	if err !=nil{
		return c.String(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func DeleteCountry(c echo.Context)error  {
	id:=c.Param("id")
	err:= countryService.DeleteCountry(id)
	if err !=nil{
		return c.String(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,"delete done")
}