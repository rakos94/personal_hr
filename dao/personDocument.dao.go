package dao

import "personal_hr/models"

type PersonDocumentDao interface {
	CreatePersonDocumnt(data * models.PersonDocument)error
	GetByIdPersonDocument(id string)(models.PersonDocument,error)
	GetByPersonId(personId string)(models.PersonDocument,error)
}
