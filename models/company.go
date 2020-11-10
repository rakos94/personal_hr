package models

// Company ...
type Company struct {
	BaseModels
	CompanyName string `gorm:"not null" json:"company_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// TableName ...
func (Company) TableName() string {
	return "tb_companies"
}
