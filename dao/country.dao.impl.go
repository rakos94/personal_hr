package dao

import "personal_hr/models"

type CountryDaoImpl struct {}

func (CountryDaoImpl)CreateCountry(data *models.Country)(*models.Country,error)  {
	result := g.Create(&data)
	if result.Error == nil {
		return data, nil
	}
	return nil, result.Error
}