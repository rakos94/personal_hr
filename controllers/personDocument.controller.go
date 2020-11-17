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