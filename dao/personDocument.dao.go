package dao

import "personal_hr/models"

type PersonDocumentDao interface {
	CreatePersonDocumnt(data * models.PersonDocument)error
}
