package dao

import "personal_hr/models"

type LocationDaoimpl struct {}

func (LocationDaoimpl)CreateLocation(data *models.Location)(*models.Location,error)  {
	result := g.Create(&data)
	if result.Error == nil {
		return data, nil
	}
	return nil, result.Error
}