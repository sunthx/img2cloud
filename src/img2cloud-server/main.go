package main

import (
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
	"encoding/json"
	"img2cloud-server/utils"
)

func main() {
	fmt.Print("img2cloud Start Running ... " + time.Now().String())

	mux := http.NewServeMux()
	mux.HandleFunc("/file/upload",uploadFile)

	server := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	server.ListenAndServe()
}

func uploadFile(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(1024*1024*1024)
	file := request.MultipartForm.File["uploaded"][0]
	fileReader,err := file.Open()
	if err != nil {
		return
	}

	data,err := ioutil.ReadAll(fileReader)
	if err != nil {
		return
	}

	fileId,err := utils.SaveFile(data)
	if err != nil {
		return
	}

	result := make(map[string]string)
	result["result"] = fileId

	value,err := json.Marshal(&result)
	if err != nil{
		return
	}

	writer.Header().Set("Content-Type","application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write(value)
}