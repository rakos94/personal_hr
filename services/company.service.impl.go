package services

import (
	"personal_hr/dao"
	"personal_hr/models"
)

var companyDao dao.CompanyDao = dao.CompanyDaoImpl{}

// CompanyServiceImpl ...
type CompanyServiceImpl struct{}

// CreateCompany ...
func (CompanyServiceImpl) CreateCompany(company *models.Company) (*models.Company, error) {
	return companyDao.CreateCompany(company)
}

// GetCompanyByID ...
func (CompanyServiceImpl) GetCompanyByID(id string) (models.Company, error) {
	return companyDao.GetCompanyByID(id)
}

// GetCompanyAll ...
func (CompanyServiceImpl) GetCompanyAll() ([]models.Company, error) {
	return companyDao.GetCompanyAll()
}

// UpdateCompany ...
func (CompanyServiceImpl) UpdateCompany(id string, company *models.Company) (*models.Company, error) {
	return companyDao.UpdateCompany(id, company)
}
