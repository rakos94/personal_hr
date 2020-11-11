package dao

import "personal_hr/models"

// PersonDao ...
type PersonDao interface {
	CreatePerson(data *models.Person) (*models.Person, error)
	GetPersonByID(id string) (models.Person, error)
	GetPersonByEmail(email string) (models.Person, error)
}
