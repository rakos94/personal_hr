package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type City struct {
	BaseModels
	Name string `gorm:"column:name" json:"name"`
	Code string	`gorm:"column:code;unique" json:"code"`
	ProviceId string `gorm:"column:provice_id" json:"provice_id"`
	Provices Provices `gorm:"foreignKey:provice_id" json:"provices"`
}
func (City) TableName() string {
	return "tb_city"
}
func (j* City)BeforeCreate(tx *gorm.DB)error  {
	id:=uuid.New()
	j.Id = id.String()
	return nil
}