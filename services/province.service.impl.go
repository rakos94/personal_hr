package services

import (
	"personal_hr/dao"
	"personal_hr/models"
)

var provinceDao dao.ProvinceDao = dao.ProvinceDaoImpl{}
var countryService CountryService = CountryServiceImpl{}

type ProvinceServiceImp struct {}

func (ProvinceServiceImp)CreateProvince(data * models.Provinces)(error)  {
	_,err := countryService.GetCountryById(data.CountryId)
	if err!= nil{
		return err
	}
	return provinceDao.CreateProvince(data)
}