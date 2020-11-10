package models

import "github.com/go-playground/validator/v10"

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
