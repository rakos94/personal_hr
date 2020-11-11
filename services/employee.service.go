package services

import "personal_hr/models"

// EmployeeService ...
type EmployeeService interface {
	CreateEmployee(emp *models.Employee) (*models.Employee, error)
	GetEmployeeByID(id string) (models.Employee, error)
	GetEmployeeAll() ([]models.Employee, error)
}
