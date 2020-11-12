package models

// Position ...
type Position struct {
	BaseModels
	Name        string  `gorm:"column:name" json:"name"`
	Code        string  `gorm:"column:code;unique" json:"code"`
	Level       int     `gorm:"column:level" json:"level"`
	Description string  `gorm:"column:description" json:"description"`
	CompanyId   string  `gorm:"column:company_id" json:"company_id"`
	Company     Company `gorm:"foreignKey:company_id" json:"company"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// TableName ...
func (Position) TableName() string {
	return "tb_positions"
}
