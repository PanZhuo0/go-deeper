package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 参数和响应需要接收方双方同步
type addParma struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type addResponse struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func main() {
	// call remote procedure by HTTP
	url := `http://localhost:8080/add`
	param := addParma{
		X: 10,
		Y: 20,
	}
	// serilization param for transmit by internet
	paramBytes, _ := json.Marshal(param) //application/json
	resp, _ := http.Post(url, "application/json", bytes.NewReader(paramBytes))
	defer resp.Body.Close()

	respBytes, _ := ioutil.ReadAll(resp.Body)
	var respData addResponse
	json.Unmarshal(respBytes, &respData)
	fmt.Println(respData.Data) //result
}
