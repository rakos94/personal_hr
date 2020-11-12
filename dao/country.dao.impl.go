package dao

import "personal_hr/models"

type CountryDaoImpl struct {}

func (CountryDaoImpl)CreateCountry(data *models.Country)(error)  {
	result := g.Create(&data)
	if result.Error == nil {
		return  nil
	}
	return  result.Error
}