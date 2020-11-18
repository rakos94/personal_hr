package services

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"personal_hr/dao"
	"personal_hr/models"
	"personal_hr/requests"
)
var personDocumentDao dao.PersonDocumentDao = dao.PersonDocumntDaoImpl{}
type PersonDocumentServiceImpl struct {}

func (PersonDocumentServiceImpl)CreatePersonDocumnt(data *requests.PersonDocumntRequest)error  {
	ada,_:= personDocumentDao.GetByPersonId(data.PersonId)
	if ada.PersonId != ""{
		return errors.New("Documnent Person Sudah ada")
	}
	dec,err:=base64.StdEncoding.DecodeString(data.File)
	if err!=nil{
		return err
	}
	path := "F:/goworkspace/src/personal_hr/filepdf"
	os.Chdir(path)
	fmt.Println(os.Getwd())
	f,err:= os.Create(data.Name+data.Filetype)
	if err!=nil{
		return err
	}
	_,err =f.Write(dec)
	if err!=nil{
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	var m models.PersonDocument
	m.PersonId=data.PersonId
	m.Description=data.Description
	m.Name = data.Name
	m.FileDocument = path+"/"+data.Name+data.Filetype
	m.FileType = data.Filetype
	err = personDocumentDao.CreatePersonDocumnt(&m)
	if err!=nil{
		return err
	}
	return nil
}
func (PersonDocumentServiceImpl)GetByIdPersonDocument(id string)(requests.PersonDocumntRequest,error)  {
	var data requests.PersonDocumntRequest
	result, err :=personDocumentDao.GetByIdPersonDocument(id)
	if err != nil {
		return  data,err
	}
	buff,err := ioutil.ReadFile(result.FileDocument)
	if err != nil {
		return  data,err
	}
	data.PersonId = result.PersonId
	data.Name =result.Name
	data.Description =result.Description
	data.Filetype =result.FileType
	data.File = base64.StdEncoding.EncodeToString(buff)
	return data,nil
}
func (PersonDocumentServiceImpl)GetByPersonId(personId string)(requests.PersonDocumntRequest,error)  {
	var data requests.PersonDocumntRequest
	result, err :=personDocumentDao.GetByPersonId(personId)
	if err != nil {
		return  data,err
	}
	buff,err := ioutil.ReadFile(result.FileDocument)
	if err != nil {
		return  data,err
	}
	data.PersonId = result.PersonId
	data.Name =result.Name
	data.Description =result.Description
	data.Filetype =result.FileType
	data.File = base64.StdEncoding.EncodeToString(buff)
	return data,nil
}