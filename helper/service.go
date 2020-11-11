package helper

import (
	"encoding/json"
	"fmt"
)

// ConvertRequest convert request to other struct (m)
func ConvertRequest(i interface{}, m interface{}) {
	byteArray, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(string(byteArray)), &m)
	if err != nil {
		fmt.Println(err)
	}
}
