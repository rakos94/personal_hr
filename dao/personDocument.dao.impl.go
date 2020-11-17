package dao

import (
	"fmt"
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