package dao

import "personal_hr/models"

type ProvinceDao interface {
	CreateProvince(data * models.Provinces)(error)
	//GetProvincesAll()([]models.Provinces,error)
	//GetProvincesById(id string)(models.Provinces,error)
	//PutProvinces(id string,data*models.Provinces)error
	//DeleteProvinces(id string)error
}
