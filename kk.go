package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Object struct {
	Data string
	Kir  string //при парсинге
}

type ParseObject struct {
	Request string `json:"request"`
}

func Decode(w http.ResponseWriter, r *http.Request) {
	//request := &ParseObject{}
	decoder := json.NewDecoder(r.Body)
	request := make(map[string]interface{})
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	str, _ := request["request"].(string)

	//JsonToData(request.Request) //request.Request содержит то что в object
	fmt.Printf(str)
}

func JsonToData(s string) {
	test := Object{}
	_ = json.Unmarshal([]byte(s), &test)

	fmt.Printf(test.Data)
	fmt.Printf(test.Kir)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/json", Decode)

	httpServer := http.Server{
		Addr:    ":6060",
		Handler: mux,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		return
	}
}
