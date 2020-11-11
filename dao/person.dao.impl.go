package dao

import (
	"personal_hr/models"
)

// PersonDaoImpl ...
type PersonDaoImpl struct{}

// CreatePerson ...
func (PersonDaoImpl) CreatePerson(person *models.Person) (*models.Person, error) {
	result := g.Create(person)
	if result.Error == nil {
		return person, nil
	}
	return nil, result.Error
}

// GetPersonByID ...
func (PersonDaoImpl) GetPersonByID(id string) (models.Person, error) {
	data := models.Person{}
	result := g.Where("id", id).First(&data)

	if result.Error == nil {
		data.Count = result.RowsAffected
		return data, nil
	}
	return data, result.Error
}

// GetPersonByEmail ...
func (PersonDaoImpl) GetPersonByEmail(email string) (models.Person, error) {
	var data models.Person
	result := g.Where("email = ?", email).Find(&data)

	if result.Error == nil {
		data.Count = result.RowsAffected
		return data, nil
	}
	return data, result.Error
}
