package dao

import "personal_hr/models"

// EmployeeDao ...
type EmployeeDao interface {
	CreateEmployee(data *models.Employee) (*models.Employee, error)
	GetEmployeeByID(id string) (models.Employee, error)
	GetEmployeeAll() ([]models.Employee, error)
	GetEmployeeCount() (int64, error)
	GetEmployeeByPersonID(personId string) (models.Employee, error)
}
