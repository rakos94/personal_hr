package models

// Person ...
type Person struct {
	BaseModels
	FirstName string  `gorm:"not null"`
	LastName  *string ``
	Email     string  `gorm:"unique;not null"`
	Password  string  `gorm:"not null"`
	Gender    *string ``
	Token     string  `gorm:"-"`
	BaseCUModels
	Count int64 `gorm:"-"`
}

// TableName ...
func (Person) TableName() string {
	return "tb_persons"
}
