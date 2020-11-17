package services

import (
	"encoding/base64"
	"fmt"
	"os"
	"personal_hr/dao"
	"personal_hr/models"
	"personal_hr/requests"
)
var personDocumentDao dao.PersonDocumentDao = dao.PersonDocumntDaoImpl{}
type PersonDocumentServiceImpl struct {}

func (PersonDocumentServiceImpl)CreatePersonDocumnt(data *requests.PersonDocumntRequest)error  {
	dec,err:=base64.StdEncoding.DecodeString(data.File)
	if err!=nil{
		return err
	}
	path := "F:/goworkspace/src/personal_hr/filepdf"
	os.Chdir(path)
	fmt.Println(os.Getwd())
	f,err:= os.Create(data.Name)
	if err!=nil{
		return err
	}
	_,err =f.Write(dec)
	if err!=nil{
		return err
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
	var m models.PersonDocument
	m.PersonId=data.PersonId
	m.Description=data.Description
	m.Name = data.Name
	m.FileDocument = path+"/"+data.Name+data.Filetype
	err = personDocumentDao.CreatePersonDocumnt(&m)
	if err!=nil{
		return err
	}
	return nil
}
