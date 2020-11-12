package models

// Assignement ...
type Assignement struct {
	BaseModels
	PositionId string   `gorm:"column:position_id;not null" json:"position_id"`
	LocationId string   `gorm:"column:location_id;not null" json:"location_id"`
	DivisionId string   `gorm:"column:division_id;not null" json:"division_id"`
	CompanyId  string   `gorm:"column:company_id" json:"company_id"`
	Position   Position `gorm:"foreignKey:position_id" json:"position"`
	Location   Location `gorm:"foreignKey:location_id" json:"location"`
	Division   Division `gorm:"foreignKey:division_id" json:"division"`
	Company    Company  `gorm:"foreignKey:company_id" json:"company"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
}

// TableName ...
func (Assignement) TableName() string {
	return "tb_assignments"
}
