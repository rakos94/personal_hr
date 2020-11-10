package models

// Position ...
type Position struct {
	BaseModels
	PositionName string `gorm:"not null" json:"position_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// TableName ...
func (Position) TableName() string {
	return "tb_positions"
}
