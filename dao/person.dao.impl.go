package dao

import (
	"personal_hr/models"
)

type PersonDaoImpl struct{}

func (PersonDaoImpl) CreatePerson(user *models.Person) (*models.Person, error) {
	result := g.Create(user)
	if result.Error == nil {
		return user, nil
	}
	return nil, result.Error
}

func (PersonDaoImpl) GetPersonById(id string) (models.Person, error) {
	data := models.Person{}
	result := g.First(&data).Where(id, id)

	if result.Error == nil {
		data.Count = result.RowsAffected
		return data, nil
	}
	return data, result.Error
}

func (PersonDaoImpl) GetPersonByEmail(email string) (models.Person, error) {
	var data models.Person
	result := g.Where("email = ?", email).Find(&data)

	if result.Error == nil {
		data.Count = result.RowsAffected
		return data, nil
	}
	return data, result.Error
}
