package models

// Employee ...
type Employee struct {
	BaseModels
	Nip      string `json:"nip" gorm:"unique;not null"`
	PersonID string `json:"person_id" gorm:"not null"`
	Person   Person `gorm:"foreignKey:PersonID"`
	BaseCUModels
}

// TableName ...
func (Employee) TableName() string {
	return "tb_employees"
}
