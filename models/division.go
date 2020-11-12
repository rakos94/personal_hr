package models

// Division ...
type Division struct {
	BaseModels
	DivisionName string     `gorm:"not null"`
	DepartmentID string     `gorm:"not null"`
	Department   Department `gorm:"foreignKey:DepartmentID"`
	BaseCUModels
}

// TableName ...
func (Division) TableName() string {
	return "tb_divisions"
}
