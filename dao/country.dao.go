package dao

import "personal_hr/models"

type CountryDao interface {
	CreateCountry(data *models.Country)(error)
	GetCountryAll()([]models.Country,error)
	GetCountryById(id string)(models.Country,error)
	PutCountry(id string,data*models.Country)error
	DeleteCountry(id string)error
}
