package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"personal_hr/configs"
	"personal_hr/models"
)

type PersonRestServiceImpl struct {}

func (PersonRestServiceImpl)GetRestPersonById(id string)(models.Person,error){
	baseURL := "http://192.168.15.245:8090/"
	url := baseURL + "api/person/"+id
	var m models.Person

	var bearer = "Bearer "+configs.ReqToken
	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response =>", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode !=200{
		return m, errors.New(string(bodyBytes))

	}

	// data := make(map[string]interface{})
	// datas := make(map[string]interface{})
	//parseData := make([]map[string]interface{}, 0, 0)

	// var resps = []interface{}{}

	json.Unmarshal(bodyBytes, &m)

	// log.Println(resps[0])
	// json.Unmarshal(resps[0], &data)


	//m.Id = parseData[0]["Id"]
	//log.Println(parseData[0]["birth_place"])
	log.Println(m)

	return m,nil
}
