package models

// Department ...
type Department struct {
	BaseModels
	DepartmentName string  `gorm:"not null"`
	CompanyID      string  `gorm:"not null"`
	Company        Company `gorm:"foreignKey:CompanyID"`
	BaseCUModels
}

// TableName ...
func (Department) TableName() string {
	return "tb_departments"
}
