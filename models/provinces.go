package models

type Provinces struct {
	BaseModels
	Name      string  `gorm:"column:name" json:"name"`
	Code      string  `gorm:"column:code;unique" json:"code"`
	CountryId string  `gorm:"column:country_id" json:"country_id"`
	Conutry   Country `gorm:"foreignKey:country_id" json:"country"`
}

func (Provinces) TableName() string {
	return "tb_provinces"
}
