package models

// Divisi ...
type Divisi struct {
	BaseModels
	DivitionName string `gorm:"not null" json:"divition_name"`
	DepartmentId string `gorm:"foreignKey:department_id;not null" json:"department_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	Department   Department
}

// TableName ...
func (Divisi) TableName() string {
	return "tb_divitions"
}
