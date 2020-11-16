package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"personal_hr/models"
	"personal_hr/services"
)

var personFamilyService services.PersonFamilyService = services.PersonFamilyServiceImpl{}

func SetPeronFamily(c *echo.Group)  {
	c.POST("/person-family",CreatePersonFamily)
	c.GET("/person-family",GetAllPesonFamily)
	c.GET("/person-family/:id",GetByIdPersonFamily)
	c.PUT("/person-family/:id",UpdatePersonFamily)
	c.DELETE("/person-family/:id",DeletePersonFamily)

}

func CreatePersonFamily(c echo.Context) error {
	data := new(models.PersonFamily)
	if err:=  c.Bind(data);err !=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	fmt.Println("controller",data)
	err := personFamilyService.CreatePersonFamily(data)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusOK,err)
}
func GetAllPesonFamily(c echo.Context)error  {
	data,err := personFamilyService.GetAllPesonFamily()
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusOK,data)
}
func GetByIdPersonFamily(c echo.Context)error  {
	id:= c.Param("id")

	data,err:=personFamilyService.GetByIdPersonFamily(id)
	if err !=nil{
		return c.String(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func UpdatePersonFamily(c echo.Context)error  {
	id:= c.Param("id")
	data := new(models.PersonFamily)
	if err:=  c.Bind(data);err !=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	err:=personFamilyService.UpdatePersonFamily(id,data)
	if err !=nil{
		return c.String(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func DeletePersonFamily(c echo.Context)error  {
	id:= c.Param("id")

	err:=personFamilyService.DeletePersonFamily(id)
	if err !=nil{
		return c.String(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,"sukses")
}