package models

// Employee ...
type Employee struct {
	BaseModels
	Nip      string `gorm:"unique;not null"`
	PersonID string `gorm:"not null"`
	Person   Person `gorm:"foreignKey:PersonID"`
	BaseCUModels
}

// TableName ...
func (Employee) TableName() string {
	return "tb_employees"
}
