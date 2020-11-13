package dao

import "personal_hr/models"

type ProvinceDaoImpl struct {}

func (ProvinceDaoImpl)CreateProvince(data * models.Provinces)( error)  {
	result := g.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
