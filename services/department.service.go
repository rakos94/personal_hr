package services

import "personal_hr/models"

// DepartmentService ...
type DepartmentService interface {
	CreateDepartment(m *models.Department) (*models.Department, error)
	GetDepartmentByID(id string) (models.Department, error)
	GetDepartmentAll() ([]models.Department, error)
	UpdateDepartment(id string, m *models.Department) (*models.Department, error)
}
