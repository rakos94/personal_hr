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
	c.GET("/province/:id/name/:name",GetProvinceByCouAndByName)
	c.GET("/province/name/:name",GetProvinceByName)
}

func CreateProvince(c echo.Context)error  {
	data  :=  new(models.Provinces)
	if err:=  c.Bind(data);err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	err := provinceService.CreateProvince(data)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,err)
}
func GetProvinceAll(c echo.Context)error  {
	data,err := provinceService.GetProvinceAll()
	if err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func GetProvinceById(c echo.Context)error  {
	id:= c.Param("id")

	data,err:=provinceService.GetProvinceById(id)
	if err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func UpdateProvince(c echo.Context)error  {
	id:= c.Param("id")
	data := new(models.Provinces)
	if err:=  c.Bind(data);err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	err:=provinceService.UpdateProvince(id,data)
	if err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func DeleteProvince(c echo.Context)error  {
	id:=c.Param("id")
	err:= provinceService.DeleteProvince(id)
	if err !=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,"delete done")
}
func GetProvinceByCouAndByName(c echo.Context)error  {
	id:= c.Param("id")
	name:= c.Param("name")
	data,err:= provinceService.GetProvinceByCouAndByName(id,name)
	if err!= nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}
func GetProvinceByName(c echo.Context) error {
	name:= c.Param("name")
	data,err:= provinceService.GetProvinceByName(name)
	if err!= nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)
}