package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"personal_hr/models"
	"personal_hr/services"
)

var provinceService services.ProvinceService = services.ProvinceServiceImpl{}
func SetProvince(c *echo.Group)  {
	c.POST("/province",CreateProvince)
	c.GET("/province",GetProvinceAll)
	c.GET("/province/:id",GetProvinceById)
	c.PUT("/province/:id",UpdateProvince)
	c.DELETE("/province/:id",DeleteProvince)
}

func CreateProvince(c echo.Context)error  {
	data  :=  new(models.Provinces)
	if err:=  c.Bind(data);err !=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	err := provinceService.CreateProvince(data)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusOK,err)
}
func GetProvinceAll(c echo.Context)error  {
	data,err := provinceService.GetProvinceAll()
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusOK,data)
}
func GetProvinceById(c echo.Context)error  {
	id:= c.Param("id")

	data,err:=provinceService.GetProvinceById(id)
	if err !=nil{
		return c.String(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func UpdateProvince(c echo.Context)error  {
	id:= c.Param("id")
	data := new(models.Provinces)
	if err:=  c.Bind(data);err !=nil{
		return c.JSON(http.StatusBadRequest,err)
	}
	err:=provinceService.UpdateProvince(id,data)
	if err !=nil{
		return c.String(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func DeleteProvince(c echo.Context)error  {
	id:=c.Param("id")
	err:= provinceService.DeleteProvince(id)
	if err !=nil{
		return c.String(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,"delete done")
}