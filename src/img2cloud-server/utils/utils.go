package utils

import (
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strings"
	"encoding/json"
	"os"
)


//生成UUID
func UUID() string {
	newGUID := uuid.Must(uuid.NewRandom())
	return newGUID.String()
}

//生成短链地址（使用百度短链API实现）
func CreateShortUrl(origin string) (string,error) {
	requestUrl := "https://dwz.cn/admin/create"

	resp, err := http.Post(requestUrl, "application/json; charset=UTF-8", strings.NewReader("{url:\""+origin+"\"}"))
	if err != nil || resp.StatusCode != 200 {
		return "",err
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)

	var result map[string]string
	err = json.Unmarshal(content,&result)
	if err != nil {
		return "",err
	}

	if val,ok := result["ShortUrl"]; ok {
		return val,nil
	}

	return "",nil
}

//SaveFile 保存文件到硬盘
func SaveFile(data []byte) (string, error) {
	fileName := UUID()
	localFilePath := "/Users/pd/Desktop/TestData/" + fileName
	newFile, err := os.Create(localFilePath)
	if err != nil {
		return "",err
	}

	defer newFile.Close()
	err = ioutil.WriteFile(localFilePath, data, os.ModePerm)
	if err != nil {
		return "",err
	}

	return fileName, nil
}
