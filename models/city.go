package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type City struct {
	BaseModels
	Name string `gorm:"column:name" json:"name"`
	Code string	`gorm:"column:code;unique" json:"code"`
	ProvinceId string `gorm:"column:province_id" json:"province_id"`
	Provinces Provinces `gorm:"foreignKey:province_id" json:"provinces"`
}
func (City) TableName() string {
	return "tb_city"
}
func (j* City)BeforeCreate(tx *gorm.DB)error  {
	id:=uuid.New()
	j.Id = id.String()
	return nil
}