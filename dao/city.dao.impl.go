package dao

import "personal_hr/models"

type CityDaoImpl struct {}

func (CityDaoImpl)CreateCity(data *models.City)(*models.City,error)  {
	result := g.Create(&data)
	if result.Error == nil {
		return data, nil
	}
	return nil, result.Error
}
