package services

import (
	"personal_hr/dao"
	"personal_hr/models"
)

var personFamilyDao dao.PersonFamilyFDao = dao.PersonFamilyDaoImpl{}
var personService PersonService =PersonServiceImpl{}
var personRestService PersonRestService = PersonRestServiceImpl{}

type PersonFamilyServiceImpl struct {}

func (PersonFamilyServiceImpl) CreatePersonFamily(data *models.PersonFamily) (error) {
	_,err := personRestService.GetRestPersonById(data.PersonId)
	if err!= nil{
		return err
	}
	return personFamilyDao.CreatePersonFamily(data)
}
func (PersonFamilyServiceImpl)GetAllPesonFamily()([]models.PersonFamily,error)  {

	return personFamilyDao.GetAllPesonFamily()
}

func (PersonFamilyServiceImpl)GetByIdPersonFamily(id string)(models.PersonFamily,error)  {

	return personFamilyDao.GetByIdPersonFamily(id)
}

func (PersonFamilyServiceImpl)UpdatePersonFamily(id string, data * models.PersonFamily) error  {
	return personFamilyDao.UpdatePersonFamily(id,data)
}
func (PersonFamilyServiceImpl)DeletePersonFamily(id string)error  {
	return personFamilyDao.DeletePersonFamily(id)
}