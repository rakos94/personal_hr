package dao

import "personal_hr/models"

type provinceDao interface {
	CreateProvince(data * models.Provinces)(*models.Provinces, error)
}
