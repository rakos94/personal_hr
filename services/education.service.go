package services

import "personal_hr/models"

// EducationService ...
type EducationService interface {
	CreateEducation(data *models.Education) (*models.Education, error)
	GetEducationAll() ([]models.Education, error)
	GetEducationByID(id string) (models.Education, error)
	UpdateEducation(id string, data *models.Education) (*models.Education, error)
}
