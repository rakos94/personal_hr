package services

import "personal_hr/models"

type PersonFamilyAddressService interface {
	CreatePersonFamilyAddress(data *models.PersonFamilyAddres)(error)
	GetAllPersonFamilyAddess()([]models.PersonFamilyAddres,error)
	GetByIdPersonFamilyAddress(id string)(models.PersonFamilyAddres,error)
	UpdatePersonFamilyAddress(id string,data *models.PersonFamilyAddres)error
	DeletePersonFamilyAddress(id string)error
}

