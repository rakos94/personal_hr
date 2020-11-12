package services

import "personal_hr/models"

type CountryService interface {
	CreateCountry(data *models.Country)(error)
}
