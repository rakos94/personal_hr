package dao

import "personal_hr/models"

// DepartmentDao ...
type DepartmentDao interface {
	CreateDepartment(data *models.Department) (*models.Department, error)
	GetDepartmentByID(id string) (models.Department, error)
	GetDepartmentAll() ([]models.Department, error)
	UpdateDepartment(id string, data *models.Department) (*models.Department, error)
	GetDepartmentByCompanyID(id string) ([]models.Department, error)
}
