package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Country struct {
	BaseModels
	Name string `gorm:"column:name" json:"name"`
	Code string	`gorm:"column:code;unique" json:"code"`
	Nasionality string `gorm:"column:nasionality" json:"nasionality"`
}
func (Country) TableName() string {
	return "tb_contry"
}
func (j* Country)BeforeCreate(tx *gorm.DB)error  {
	id:=uuid.New()
	j.Id = id.String()
	return nil
}