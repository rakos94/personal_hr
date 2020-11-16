package services

import "personal_hr/models"

type ProvinceService interface {
	CreateProvince(data * models.Provinces)(error)
	GetProvinceAll()([]models.Provinces,error)
	GetProvinceById(id string)(models.Provinces,error)
	UpdateProvince(id string,data*models.Provinces)error
	DeleteProvince(id string)error
}
