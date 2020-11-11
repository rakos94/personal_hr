package models

// Division ...
type Division struct {
	BaseModels
	DivisionName string `gorm:"not null" json:"division_name"`
	DepartmentId string `gorm:"foreignKey:department_id;not null" json:"department_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	Department   Department
}

// TableName ...
func (Division) TableName() string {
	return "tb_divisions"
}
