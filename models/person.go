package models

import (
	"database/sql"
)

// Person ...
type Person struct {
	BaseModels
	FirstName string         `gorm:"not null" form:"first_name" json:"first_name" validate:"required"`
	LastName  sql.NullString `form:"last_name" json:"last_name"`
	Email     string         `gorm:"unique;not null" form:"email" json:"email" validate:"required,email"`
	Password  string         `gorm:"not null" json:"-"`
	Gender    sql.NullString `json:"gender"`
	Token     string         `json:"token" gorm:"-"`
	BaseCUModels
	Count int64 `gorm:"-"`
}

// TableName ...
func (Person) TableName() string {
	return "tb_persons"
}
