package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type addParma struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type addResponse struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func add(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/add", addHandler)
	http.ListenAndServe(":8080", nil)
}

// server process function
func addHandler(w http.ResponseWriter, r *http.Request) {
	// analyze parameter
	b, _ := ioutil.ReadAll(r.Body)
	param := &addParma{}
	json.Unmarshal(b, param)
	// server logic
	ret := add(param.X, param.Y)
	// response
	respBytes, _ := json.Marshal(addResponse{Code: 0, Data: ret})
	w.Write(respBytes)
}
