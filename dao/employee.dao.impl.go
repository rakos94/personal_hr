package dao

import "personal_hr/models"

// EmployeeDaoImpl ...
type EmployeeDaoImpl struct {
}

// CreateEmployee ...
func (EmployeeDaoImpl) CreateEmployee(data *models.Employee) (*models.Employee, error) {
	result := g.Create(&data)
	if result.Error == nil {
		return data, nil
	}
	return nil, result.Error
}

// GetEmployeeByID ...
func (EmployeeDaoImpl) GetEmployeeByID(id string) (models.Employee, error) {
	m := models.Employee{}
	result := g.Preload("Person").Where("id", id).First(&m)
	if result.Error != nil {
		return m, result.Error
	}

	return m, nil
}

// GetEmployeeAll ...
func (EmployeeDaoImpl) GetEmployeeAll() ([]models.Employee, error) {
	m := []models.Employee{}
	result := g.Preload("Person").Find(&m)
	if result.Error != nil {
		return m, result.Error
	}

	return m, nil
}

// GetEmployeeCount ...
func (EmployeeDaoImpl) GetEmployeeCount() (int64, error) {
	var count int64
	result := g.Model(&models.Employee{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

// GetEmployeeByPersonID ...
func (EmployeeDaoImpl) GetEmployeeByPersonID(personID string) (models.Employee, error) {
	em := models.Employee{}
	result := g.Where("person_id", personID).First(&em)
	if result.Error != nil {
		return em, result.Error
	}

	return em, nil
}
