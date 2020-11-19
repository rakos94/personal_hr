package controllers

import (
	"github.com/labstack/echo"
	"personal_hr/models"
	"net/http"
	"personal_hr/services"
)
var cityService services.CityService = services.CityServiceImpl{}
func SetCity(c *echo.Group)  {
	c.POST("/city",CreateCity)
	c.GET("/city",GetCityAll)
	c.GET("/city/:id",GetCityById)
	c.PUT("/city/:id",UpdateCity)
	c.DELETE("/city/:id",DeleteCity)
	c.GET("/city/:id/name/:name",GetCityByProvAndByName)
	c.GET("/city/name/:name",GetCityByName)
}
func CreateCity(c echo.Context)error  {
	m  := new(models.City)
	if err:=  c.Bind(m);err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	err := cityService.CreateCity(m)
	if err != nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return  c.JSON(http.StatusOK," Create sukses")
}
func GetCityAll(c echo.Context)error  {
	data,err := cityService.GetCityAll()
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusOK,data)
}
func GetCityById(c echo.Context)error  {
	id:= c.Param("id")

	data,err:=cityService.GetCityById(id)
	if err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func UpdateCity(c echo.Context)error  {
	id:= c.Param("id")
	data := new(models.City)
	if err:=  c.Bind(data);err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	err:=cityService.UpdateCity(id,data)
	if err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func DeleteCity(c echo.Context)error  {
	id:=c.Param("id")
	err:= cityService.DeleteCity(id)
	if err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,"delete done")
}
func GetCityByProvAndByName(c echo.Context)error{
	id:= c.Param("id")
	name:= c.Param("name")
	data,err:= cityService.GetCityByProvAndByName(id,name)
	if err!= nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func GetCityByName(c echo.Context)error{
	name:= c.Param("name")
	data,err:= cityService.GetCityByName(name)
	if err!= nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}