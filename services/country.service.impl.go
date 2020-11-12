package services

import (
	"personal_hr/dao"
	"personal_hr/models"
)

var countryDao dao.CountryDao = dao.CountryDaoImpl{}
type CountryServiceImpl struct {}
func (CountryServiceImpl)CreateCountry(data *models.Country)(error)  {
	return countryDao.CreateCountry(data)
}
