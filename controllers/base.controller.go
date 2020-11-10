package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func SetInit(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Rest API started")
	})
}

func res(c echo.Context, data interface{}) error {
	// result, err := json.Marshal(data)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }
	return c.JSON(http.StatusOK, data)
}

func resErr(c echo.Context, err error) error {
	// result, err := json.Marshal(err.Error())
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }
	// return c.JSON(http.StatusBadRequest, string(result))
	return c.JSON(http.StatusBadRequest, err.Error())
}

func resValErr(c echo.Context, err error) error {
	var errorMessages []string
	validationErrors := err.(validator.ValidationErrors)
	for _, v := range validationErrors {
		var message string
		switch v.Tag() {
		case "required":
			message = v.Field() + " is required"
		case "email":
			message = v.Field() + " must be an email"
		}
		errorMessages = append(errorMessages, message)
	}
	return c.JSON(http.StatusBadRequest, errorMessages)
}
