package services

import (
	"personal_hr/requests"
)

type PersonDocumentService interface {
	CreatePersonDocumnt(data *requests.PersonDocumntRequest)error
	GetByIdPersonDocument(id string)(requests.PersonDocumntRequest,error)
	GetByPersonId(personId string)(requests.PersonDocumntRequest,error)
}