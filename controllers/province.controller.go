package controllers

import (
	"personal_hr/services"
)

var provinceservice services.ProvinceService = services.ProvinceServiceImp{}
//func SetProvince(c *echo.Group){
//	c.POST("/")
//}
//
//func CreateProvince(e * echo.Context)error  {
//
//}