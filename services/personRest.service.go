package services

import "personal_hr/models"

type PersonRestService interface {
	GetRestPersonById(id string)(models.Person,error)
}
