package services

import "personal_hr/models"

// CompanyService ...
type CompanyService interface {
	CreateCompany(company *models.Company) (*models.Company, error)
	GetCompanyByID(id string) (models.Company, error)
	GetCompanyAll() ([]models.Company, error)
	UpdateCompany(id string, company *models.Company) (models.Company, error)
}
