package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"personal_hr/models"
	"personal_hr/services"
)
var personFamilyAddressService services.PersonFamilyAddressService = services.PersonFamilyAddressServiceImpl{}

func SetPeronFamilyAddress(c *echo.Group)  {
	c.POST("/person-family-address",CreatePersonFamilyAddress)
	c.GET("/person-family-address",GetAllPersonFamilyAddess)
	c.GET("/person-family-address/:id",GetByIdPersonFamilyAddress)
	c.PUT("/person-family-address/:id",UpdatePersonFamilyAddress)
	c.DELETE("/person-family-address/:id",DeletePersonFamilyAddress)
}
func CreatePersonFamilyAddress(c echo.Context)error  {
	data := new(models.PersonFamilyAddres)
	if err:=  c.Bind(data);err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	err := personFamilyAddressService.CreatePersonFamilyAddress(data)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,err)
}
func GetAllPersonFamilyAddess(c echo.Context)error  {
	data,err := personFamilyAddressService.GetAllPersonFamilyAddess()
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func GetByIdPersonFamilyAddress(c echo.Context)error  {
	id:= c.Param("id")

	data,err:=personFamilyAddressService.GetByIdPersonFamilyAddress(id)
	if err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func UpdatePersonFamilyAddress(c echo.Context)error  {
	id:= c.Param("id")
	data := new(models.PersonFamilyAddres)
	if err:=  c.Bind(data);err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	err:=personFamilyAddressService.UpdatePersonFamilyAddress(id,data)
	if err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func DeletePersonFamilyAddress(c echo.Context)error  {
	id:= c.Param("id")

	err:=personFamilyService.DeletePersonFamily(id)
	if err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,"sukses")
}