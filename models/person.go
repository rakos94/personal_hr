package models

import (
	"personal_hr/requests"

	"gorm.io/datatypes"
)

// Person ...
type Person struct {
	BaseModels
	FirstName         string         `gorm:"not null"`
	LastName          *string        ``
	Email             string         `gorm:"unique;not null"`
	Password          string         `gorm:"not null"`
	BirthPlace        *string        ``
	BirthDate         datatypes.Date `gorm:"not null"`
	Mobile            *string        ``
	Gender            string         `gorm:"not null"`
	IsFromRecruitment bool           `gorm:"not null;default:false"`
	CreatedBy         string         `gorm:"not null;default:admin"`
	UpdatedBy         string         `gorm:"not null;default:admin"`
	Token             string         `gorm:"-"`
	BaseCUModels
}

// TableName ...
func (Person) TableName() string {
	return "cor_persons"
}

// ConvertFromRequest ...
func (Person) ConvertFromRequest(r *requests.PersonUpdateRequest) *Person {
	return &Person{
		FirstName:         r.FirstName,
		LastName:          r.LastName,
		Email:             r.Email,
		BirthPlace:        r.BirthPlace,
		BirthDate:         datatypes.Date(r.BirthDate),
		Mobile:            r.Mobile,
		Gender:            r.Gender,
		IsFromRecruitment: r.IsFromRecruitment,
	}
}
