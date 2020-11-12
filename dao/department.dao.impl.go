package dao

import "personal_hr/models"

// DepartmentDaoImpl ...
type DepartmentDaoImpl struct{}

// CreateDepartment ...
func (DepartmentDaoImpl) CreateDepartment(data *models.Department) (*models.Department, error) {
	result := g.Create(data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// GetDepartmentByID ...
func (DepartmentDaoImpl) GetDepartmentByID(id string) (models.Department, error) {
	m := models.Department{}
	result := g.Where("id", id).First(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

// GetDepartmentAll ...
func (DepartmentDaoImpl) GetDepartmentAll() ([]models.Department, error) {
	m := []models.Department{}
	result := g.Find(&m)
	if result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

// UpdateDepartment ...
func (DepartmentDaoImpl) UpdateDepartment(id string, data *models.Department) (*models.Department, error) {
	m := &models.Department{}
	update := g.Where("id", id).First(m).Updates(data)
	if update.Error != nil {
		return m, update.Error
	}

	return m, nil
}

// GetDepartmentByCompanyID ...
func (DepartmentDaoImpl) GetDepartmentByCompanyID(id string) ([]models.Department, error) {
	m := []models.Department{}
	result := g.Where("company_id", id).Find(&m)
	if result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}
