package dao

import "personal_hr/models"

type CountryDao interface {
	CreateCountry(data *models.Country)(error)
}
