package services

import "personal_hr/models"

// PersonService ...
type PersonService interface {
	CreatePerson(person *models.Person) (*models.Person, error)
	GetPersonByID(id string) (models.Person, error)
	Login(email string, pwd string) (models.Person, error)
	UpdatePerson(id string, person *models.Person) (*models.Person, error)
	UpdatePassword(id string, password string) error
}
