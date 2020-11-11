package dao

import "personal_hr/models"

type locationDao interface {
	CreateLocation(data *models.Location)(*models.Location,error)
}
