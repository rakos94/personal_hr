package models

// Location ...
type Location struct {
	BaseModels
	LocationName string `gorm:"not null" json:"Location_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// TableName ...
func (Location) TableName() string {
	return "tb_locations"
}
