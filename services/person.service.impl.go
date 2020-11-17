package services

import (
	"errors"
	"log"
	"personal_hr/configs"
	"personal_hr/dao"
	"personal_hr/grpc"
	"personal_hr/models"
	pb "personal_hr/models"

	"golang.org/x/crypto/bcrypt"
)

var personDao dao.PersonDao = dao.PersonDaoImpl{}

// PersonServiceImpl ...
type PersonServiceImpl struct{}

// Login ...
func (PersonServiceImpl) Login(email string, pwd string) (models.Person, error) {
	result, err := personDao.GetPersonByEmail(email)
	if err != nil {
		return models.Person{}, errors.New("Username/password not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(pwd))
	if err != nil {
		log.Println(err.Error())
		return models.Person{}, err
	}

	result.Token, err = configs.CreateJwtToken(result.Email)
	if err != nil {
		log.Println(err.Error())
		return models.Person{}, err
	}

	return result, nil
}

// CreatePerson ...
func (PersonServiceImpl) CreatePerson(person *models.Person) (*models.Person, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(person.Password), 4)
	if err != nil {
		return nil, err
	}

	person.Password = string(result)
	rs, err := personDao.CreatePerson(person)
	if err != nil {
		return nil, err
	}

	// defer grpc.Conn.Close())
	res, err := grpc.Client.Register(grpc.Ctx,
		&pb.Users{Username: person.Email, Password: person.Password})
	if err != nil {
		log.Println("Error credential register =>", err)
		return nil, err
	}

	log.Println("Credential created ->", res)

	return rs, nil
}

// GetPersonAll ...
func (PersonServiceImpl) GetPersonAll() ([]models.Person, error) {
	return personDao.GetPersonAll()
}

// GetPersonByID ...
func (PersonServiceImpl) GetPersonByID(id string) (models.Person, error) {
	return personDao.GetPersonByID(id)
}

// UpdatePerson ...
func (PersonServiceImpl) UpdatePerson(id string, person *models.Person) (*models.Person, error) {
	return personDao.UpdatePerson(id, person)
}

// UpdatePassword ...
func (s PersonServiceImpl) UpdatePassword(id string, password string) error {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return err
	}

	return personDao.UpdatePassword(id, string(result))
}
