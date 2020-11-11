package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Provinces struct {
	BaseModels
	Name string `gorm:"column:name" json:"name"`
	Code string	`gorm:"column:code;unique" json:"code"`
	CountryId string `gorm:"column:country_id" json:"country_id"`
	Conutry Country `gorm:"foreignKey:country_id" json:"country"`

}
func (Provinces) TableName() string {
	return "tb_provicesy"
}
func (j* Provinces)BeforeCreate(tx *gorm.DB)error  {
	id:=uuid.New()
	j.Id = id.String()
	return nil
}