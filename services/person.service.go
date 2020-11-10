package services

import "personal_hr/models"

type PersonService interface {
	CreatePerson(Person *models.Person) (*models.Person, error)
	GetPersonById(id string) (models.Person, error)
	Login(email string, pwd string) (models.Person, error)
}
