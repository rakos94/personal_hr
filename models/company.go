package models

// Company ...
type Company struct {
	BaseModels
	CompanyName string `gorm:"not null" `
	BaseCUModels
}

// TableName ...
func (Company) TableName() string {
	return "tb_companies"
}
