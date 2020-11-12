package dao

import (
	"errors"
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
	if result.RowsAffected == 0 {
		return data, errors.New("Record not found")
	}
	if result.Error != nil {
		return data, result.Error
	}

	data.Count = result.RowsAffected
	return data, nil
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

// UpdatePerson ...
func (PersonDaoImpl) UpdatePerson(id string, data *models.Person) (*models.Person, error) {
	m := &models.Person{}
	first := g.Where("id", id).First(m)
	if first.Error != nil {
		return m, first.Error
	}
	update := g.Model(m).Updates(data)
	if update.Error != nil {
		return m, update.Error
	}

	return m, nil
}

// UpdatePassword ...
func (PersonDaoImpl) UpdatePassword(id string, password string) error {
	m := &models.Person{}
	update := g.Model(m).Where("id", id).Update("password", password)
	if update.Error != nil {
		return update.Error
	}

	return nil
}
