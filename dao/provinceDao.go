package dao

import "personal_hr/models"

type ProvinceDao interface {
	CreateProvince(data * models.Provinces)(error)
	GetProvinceAll()([]models.Provinces,error)
	GetProvinceById(id string)(models.Provinces,error)
	UpdateProvince(id string,data*models.Provinces)error
	DeleteProvince(id string)error
	GetProvinceByCouAndByName(id string, name string)([]models.Provinces,error)
	GetProvinceByName(name string)([]models.Provinces,error)
}
