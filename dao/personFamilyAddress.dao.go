package dao

import "personal_hr/models"

type PersonFamilyAddress interface {
	CreatePersonFamilyAddress(data *models.PersonFamilyAddres)(error)
	GetAllPersonFamilyAddess()([]models.PersonFamilyAddres,error)
	GetByIdPersonFamilyAddress(id string)(models.PersonFamilyAddres,error)
	UpdatePersonFamilyAddress(id string,data *models.PersonFamilyAddres)error
	DeletePersonFamilyAddress(id string)error
}
