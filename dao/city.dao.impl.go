package dao

import "personal_hr/models"

type CityDaoImpl struct {}

func (CityDaoImpl)CreateCity(data *models.City)(error)  {
	result := g.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (CityDaoImpl)GetCityAll()([]models.City,error)  {
	var  c []models.City
	result :=  g.Preload("Provinces.Country").Find(&c)
	if result.Error !=nil{
		return  nil,result.Error
	}
	return c,nil
}
func (CityDaoImpl)GetCityById(id string)(models.City,error)  {
	var c models.City
	result := g.Preload("Provinces.Country").Where("id",id).First(&c)

	return c,result.Error
}
func (CityDaoImpl)UpdateCity(id string,data*models.City)error  {
	var m = new(models.City)
	first := g.Where("id",id).First(&m)
	//subq:= g.Select("id").Where("name",id).Find(&models.Provinces{})
	//q:= g.Where("province_id = (?)",subq).Find(&m)

	if first.Error != nil {
		return first.Error
	}
	update := g.Model(m).Updates(data)
	if update.Error != nil {
		return update.Error
	}

	return nil
}
func (CityDaoImpl)DeleteCity(id string)error  {
	var m = new(models.City)
	first := g.Where("id",id).Delete(&m)
	if first.Error != nil {
		return first.Error
	}
	return nil
}
func (CityDaoImpl)GetCityByProvAndByName(id string, name string)([]models.City,error)  {
	var(
		m []models.City
	)
	data := g.Preload("Provinces.Country").Where("province_id = ? AND name Like ?", id,"%"+name+"%").Find(&m)
	if data.Error != nil{
		return m,data.Error
	}
	return m,nil
}
func (CityDaoImpl)GetCityByName(name string)([]models.City,error)  {
	var(
		m []models.City
	)
	data := g.Preload("Provinces.Country").Where("name Like ?", "%"+name+"%").Find(&m)
	if data.Error != nil{
		return m,data.Error
	}
	return m,nil
}