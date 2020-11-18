package pb_client

import (
	"log"
	"personal_hr/configs"
	pbx "personal_hr/pb.client/model"
)


func GetPersonById(id string)(pbx.GetPerson,error)   {
	var m = pbx.GetPerson{
		Id:id,
	}
	log.Println("test id :",id)
	res,err:=configs.Clients.GetPersonById(configs.CtxClient,&m)
	if err!=nil{
		log.Println("masuk ini : ",err)
		return m,err
	}
	m.Id = res.Data


	return m,nil

}
