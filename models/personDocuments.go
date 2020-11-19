package models

type PersonDocument struct {
	BaseModels
	Name string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	FileDocument string `gorm:"column:file_document" json:"file_document"`
	FileType string `gorm:"column:file_type" json:"file_type"`
	PersonId string `gorm:"column:person_id;unique" json:"person_id"`
	Person Person `gorm:"foreignKey:person_id" json:"person"`
}

func (PersonDocument) TableName() string {
	return "cor_person_document"
}
