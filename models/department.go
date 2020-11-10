package models

// Department ...
type Department struct {
	BaseModels
	DepartmentName string `gorm:"not null" json:"department_name"`
	CompanyId      string `gorm:"foreignKey:company_id;not null" json:"company_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Company        Company
}

// TableName ...
func (Department) TableName() string {
	return "tb_departments"
}
