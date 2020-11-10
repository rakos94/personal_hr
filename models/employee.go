package models

// Employee ...
type Employee struct {
	BaseModels
	Nip      string `gorm:"unique;not null" json:"nip"`
	PersonId string `gorm:"foreignKey:person_id;not null" json:"person_id"`
	Person   Person
	BaseCUModels
}

// TableName ...
func (Employee) TableName() string {
	return "tb_employees"
}
