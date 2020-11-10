package models

// Assignement ...
type Assignement struct {
	BaseModels
	PositionId string `gorm:"foreignKey:position_id;not null" json:"position_id"`
	LocationId string `gorm:"foreignKey:location_id;not null" json:"location_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Position   Position
	Location   Location
}

// TableName ...
func (Assignement) TableName() string {
	return "tb_assignements"
}
