package main

import (
	"fmt"
	"personal_hr/configs"
	"personal_hr/controllers"
	"personal_hr/dao"
	"personal_hr/grpc"
	"personal_hr/models"
	"personal_hr/services"

	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	g := NewConn()
	dao.SetDao(g)
	services.SetService(g)

	//set jwt
	jwtGroup := configs.SetJwt(e)

	// set validator
	e.Validator = &models.CustomValidator{Validator: validator.New()}

	grpc.ClientConnect()

	controllers.SetInit(e)
	controllers.SetPerson(jwtGroup, e)
	controllers.SetCountry(jwtGroup)
	controllers.SetPeronFamily(jwtGroup)
	controllers.SetPeronFamilyAddress(jwtGroup)
	controllers.SetProvince(jwtGroup)
	controllers.SetCity(jwtGroup)
	controllers.SetEducation(jwtGroup)
	controllers.SetAddress(jwtGroup)

	//start server
	e.Logger.Fatal(e.Start(":1234"))
}

// NewConn ...
func NewConn() *gorm.DB {
	g, err := configs.Conn()
	if err == nil {
		configs.MigrateSchema(g)
		fmt.Println("Success Create Table")
	}

	return g
}
