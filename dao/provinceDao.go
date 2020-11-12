package dao

import "personal_hr/models"

type ProvinceDao interface {
	CreateProvince(data * models.Provinces)(*models.Provinces, error)
	GetCountryAll()([]models.Country,error)
	GetCountryById(id string)(models.Country,error)
	PutCountry(id string,data*models.Country)error
	DeleteCountry(id string)error
}
