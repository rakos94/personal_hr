package dao

import (
	"personal_hr/models"
)

// CompanyDaoImpl ...
type CompanyDaoImpl struct {
}

// CreateCompany ...
func (CompanyDaoImpl) CreateCompany(data *models.Company) (*models.Company, error) {
	result := g.Create(data)
	if result.Error == nil {
		return data, nil
	}
	return nil, result.Error
}

// GetCompanyByID ...
func (CompanyDaoImpl) GetCompanyByID(id string) (models.Company, error) {
	m := models.Company{}
	result := g.Where("id", id).First(&m)
	if result.Error != nil {
		return m, result.Error
	}

	return m, nil
}

// GetCompanyAll ...
func (CompanyDaoImpl) GetCompanyAll() ([]models.Company, error) {
	m := []models.Company{}
	result := g.Find(&m)
	if result.Error != nil {
		return m, result.Error
	}

	return m, nil
}

// UpdateCompany ...
func (CompanyDaoImpl) UpdateCompany(id string, data *models.Company) (models.Company, error) {
	m := models.Company{}
	first := g.Where("id", id).First(&m)
	if first.Error != nil {
		return m, first.Error
	}
	update := g.Model(&m).Updates(data)
	if update.Error != nil {
		return m, update.Error
	}

	return m, nil
}
