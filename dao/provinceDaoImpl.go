package dao

import "personal_hr/models"

type ProvinceDaoImpl struct {}

func (ProvinceDaoImpl)CreateProvince(data * models.Provinces)(*models.Provinces, error)  {
	result := g.Create(&data)
	if result.Error == nil {
		return data, nil
	}
	return nil, result.Error
}
