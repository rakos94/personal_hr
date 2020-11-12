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

func (CountryServiceImpl)GetCountryAll()([]models.Country,error)  {
	return countryDao.GetCountryAll()
}
func (CountryServiceImpl)GetCountryById(id string)(models.Country,error)   {
	return countryDao.GetCountryById(id)
}
func (CountryServiceImpl)PutCountry(id string,data*models.Country)error  {
	_,err:= countryDao.GetCountryById(id)
	if err!=nil {
		return err
	}
	return countryDao.PutCountry(id,data)
}
func (CountryServiceImpl)DeleteCountry(id string)error  {
	return countryDao.DeleteCountry(id)
}