package services

import (
	"errors"
	"log"
	"personal_hr/dao"
	"personal_hr/models"

	"golang.org/x/crypto/bcrypt"
)

var personDao dao.PersonDao = dao.PersonDaoImpl{}

// PersonServiceImpl ...
type PersonServiceImpl struct{}

// CreatePerson ...
func (PersonServiceImpl) CreatePerson(person *models.Person) (*models.Person, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(person.Password), 4)
	if err == nil {
		person.Password = string(result)
		return personDao.CreatePerson(person)
	}
	return nil, err
}

// GetPersonById ...
func (PersonServiceImpl) GetPersonById(id string) (models.Person, error) {
	return personDao.GetPersonByID(id)
}

// Login ...
func (PersonServiceImpl) Login(email string, pwd string) (models.Person, error) {
	result, err := personDao.GetPersonByEmail(email)
	if err == nil && result.Count > 0 {
		var err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(pwd))
		if err != nil {
			log.Println(err.Error())
		}
		if err == nil {
			return result, nil
		}
	}

	return models.Person{}, errors.New("Username/password not found")
}
