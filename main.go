// PaperManagementClient project main.go
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/robfig/cron"
)

const (
	//主机地址
	HTTP_HOST string = "http://180.76.163.58:8081/factory"
	//	HTTP_HOST        string = "http://localhost:8081/factory"
	HTTP_APPLICATION string = "application/json;charset=utf-8"
	FILE_NAME        string = "infor.txt"
)

func main() {
	c := cron.New()
	//秒 分 时 日 月 星期
	//	spec := "0 */1 * * * *" //每分钟一次
	spec := "*/10 * * * * *"
	c.AddFunc(spec, func() {

		//读取文件
		inputFile := FILE_NAME
		buf, err := ioutil.ReadFile(inputFile)
		if err != nil {
			log.Println(os.Stderr, "File Error: %s\n", err)
		}
		fileContent := string(buf)
		log.Println("file content:\n", fileContent)

		//解析内容
		//		paperManageMap := ParseFileContent(fileContent)
		//		log.Println("paperManageMap:\n", paperManageMap)
		if !checkData(fileContent) {
			return
		}

		//解析完成后通过http协议发送到服务端
		httpPost(fileContent)

	})
	c.Start()
	select {} //阻塞主线程不退出
}

/*
解析文件内容
*/
func ParseFileContent(fileContent string) (paperManageMap map[string]string) {
	paperManageMap = make(map[string]string)
	fileContentSlices := strings.Split(fileContent, ",")
	//	log.Println("fileContentSlices:\n", fileContentSlices)
	if len(fileContentSlices) == 0 {
		log.Println("fileContentSlices length is 0")
		return
	}
	for _, value := range fileContentSlices {
		if len(value) > 0 {
			valueSlices := strings.Split(value, ":")
			if len(valueSlices) == 0 {
				log.Println("valueSlices length is 0")
				return
			}
			paperManageMap[valueSlices[0]] = valueSlices[1]
		}
	}
	return paperManageMap
}

/*
判断是否为json格式
*/
func checkData(data string) bool {
	type CheckJson struct {
		Factory string `json:"Factory"`
		Other   string `json:"Other"`
	}
	var checkJson CheckJson
	err := json.Unmarshal([]byte(data), &checkJson)
	if err != nil {
		log.Println("json data invalid")
		return false
	}
	return true
}

/*
发送到服务端
*/
func httpPost(data string) {
	body := bytes.NewBuffer([]byte(data))
	resp, err := http.Post(HTTP_HOST, HTTP_APPLICATION, body)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	log.Println(string(respBody))

}
