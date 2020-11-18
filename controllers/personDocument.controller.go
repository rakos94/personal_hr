package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"personal_hr/requests"
	"personal_hr/services"
)

var personDocumentService services.PersonDocumentService = services.PersonDocumentServiceImpl{}

func SetPersonDocument(e *echo.Group)  {
	e.POST("/person-document",CreatePersonDocumnt)
	e.GET("/person-document/:id",GetByIdPersonDocument)
}

func CreatePersonDocumnt(c echo.Context)error  {
	data := new(requests.PersonDocumntRequest)
	if err := c.Bind(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err := personDocumentService.CreatePersonDocumnt(data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK,err)
}
func GetByIdPersonDocument(c echo.Context)error  {
	id:= c.Param("id")
	data,err:=personDocumentService.GetByIdPersonDocument(id)
	if err !=nil{
		return c.String(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,data)

}