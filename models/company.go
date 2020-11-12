package models

// Company ...
type Company struct {
	BaseModels
	CompanyName string       `gorm:"not null"`
	Department  []Department `gorm:"foreignKey:CompanyID"`
	BaseCUModels
}

// TableName ...
func (Company) TableName() string {
	return "tb_companies"
}
