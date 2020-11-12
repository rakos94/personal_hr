package dao

import "personal_hr/models"

// CompanyDao ...
type CompanyDao interface {
	CreateCompany(data *models.Company) (*models.Company, error)
	GetCompanyByID(id string) (models.Company, error)
	GetCompanyAll() ([]models.Company, error)
	UpdateCompany(id string, data *models.Company) (*models.Company, error)
}
