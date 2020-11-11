package models

import (
	"database/sql"
)

// Person ...
type Person struct {
	BaseModels
	FirstName string         `gorm:"not null"`
	LastName  sql.NullString ``
	Email     string         `gorm:"unique;not null"`
	Password  string         `gorm:"not null"`
	Gender    sql.NullString ``
	Token     string         `gorm:"-"`
	BaseCUModels
	Count int64 `gorm:"-"`
}

// TableName ...
func (Person) TableName() string {
	return "tb_persons"
}
