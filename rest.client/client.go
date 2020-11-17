package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// jsonReq, err := json.Marshal(Request{
	// 	Username: "test@gmail.com",
	// 	Password: "123456",
	// })
	baseURL := "http://192.168.15.245:1234/"
	url := baseURL + "api/person"

	var bearer = "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImthbWFsMUBhZG1pbi5jb20iLCJleHAiOjE2MDU2MjIyNDN9.e3L5UV68-J-_GAjUuupQSLn4yAzm22OCXmpX42g-Dsq4Uik1KGg0cKATyTbkaBmNteCkhmPS7UMV-_LsLOZ5mQ"

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
	log.Println(string(bodyBytes))

	// data := make(map[string]interface{})
	// datas := make(map[string]interface{})
	parseData := make([]map[string]interface{}, 0, 0)

	// var resps = []interface{}{}

	json.Unmarshal(bodyBytes, &parseData)

	// log.Println(resps[0])
	// json.Unmarshal(resps[0], &data)

	log.Println(parseData[0]["birth_place"])
}
