package dao

import (
	"fmt"
	"log"
	"personal_hr/models"
)

type PersonDocumntDaoImpl struct {}

func (PersonDocumntDaoImpl)CreatePersonDocumnt(data * models.PersonDocument)error  {
	fmt.Println(data)
	result := g.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (PersonDocumntDaoImpl)GetByIdPersonDocument(id string)(models.PersonDocument,error) {
	data := models.PersonDocument{}
	result := g.Where("id", id).First(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}
func (PersonDocumntDaoImpl)GetByPersonId(personId string)(models.PersonDocument,error)  {
	data := models.PersonDocument{}
	log.Println("ip",personId)
	result := g.Where("person_id = (?)", personId).First(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}