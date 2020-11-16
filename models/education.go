package models

// Education ...
type Education struct {
	BaseModels
	PersonID    string   `gorm:"not null"`
	Institution string   `gorm:"not null"`
	Subject     string   `gorm:"not null"`
	Grade       *float32 `gorm:"not null"`
	YearBegin   *int     ``
	YearEnd     *int     ``
	CreatedBy   string   `gorm:"not null;default:admin"`
	UpdatedBy   string   `gorm:"not null;default:admin"`
	Person      Person   `gorm:"foreignKey:PersonID"`
	BaseCUModels
}

// TableName ...
func (Education) TableName() string {
	return "cor_person_educations"
}

// // ConvertFromRequest ...
// func (Person) ConvertFromRequest(r *requests.PersonUpdateRequest) *Person {
// 	return &Person{
// 		FirstName:         r.FirstName,
// 		LastName:          r.LastName,
// 		Email:             r.Email,
// 		BirthPlace:        r.BirthPlace,
// 		BirthDate:         datatypes.Date(r.BirthDate),
// 		Mobile:            r.Mobile,
// 		Gender:            r.Gender,
// 		IsFromRecruitment: r.IsFromRecruitment,
// 	}
// }
