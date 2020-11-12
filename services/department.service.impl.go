package services

import (
	"errors"
	"personal_hr/dao"
	"personal_hr/models"
)

var departmentDao dao.DepartmentDao = dao.DepartmentDaoImpl{}
var companyService CompanyService = CompanyServiceImpl{}

// DepartmentServiceImpl ...
type DepartmentServiceImpl struct{}

// CreateDepartment ...
func (DepartmentServiceImpl) CreateDepartment(m *models.Department) (*models.Department, error) {
	// check company id exist
	_, err := companyService.GetCompanyByID(m.CompanyID)
	if err != nil {
		return nil, errors.New("Company id not found")
	}
	return departmentDao.CreateDepartment(m)
}

// GetDepartmentByID ...
func (DepartmentServiceImpl) GetDepartmentByID(id string) (models.Department, error) {
	return departmentDao.GetDepartmentByID(id)
}

// GetDepartmentAll ...
func (DepartmentServiceImpl) GetDepartmentAll() ([]models.Department, error) {
	return departmentDao.GetDepartmentAll()
}

// UpdateDepartment ...
func (DepartmentServiceImpl) UpdateDepartment(id string, m *models.Department) (*models.Department, error) {
	return departmentDao.UpdateDepartment(id, m)
}

// GetDepartmentByCompanyID ...
func (DepartmentServiceImpl) GetDepartmentByCompanyID(id string) ([]models.Department, error) {
	return departmentDao.GetDepartmentByCompanyID(id)
}
