package dao

import "personal_hr/models"

type CityDao interface {
	CreateCity(data *models.City)(*models.City,error)
}
