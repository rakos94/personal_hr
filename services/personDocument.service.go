package services

import "personal_hr/requests"

type PersonDocumentService interface {
	CreatePersonDocumnt(data *requests.PersonDocumntRequest)error
}