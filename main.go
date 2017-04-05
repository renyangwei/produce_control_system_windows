// PaperManagementClient project main.go
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
		fd, err := os.Open(FILE_NAME)
		if err != nil {
			log.Println("open file err:", err.Error())
		}
		br := bufio.NewReader(fd)
		r, _, err := br.ReadRune()
		if err != nil {
			log.Println("read rune err:", err.Error())
		}
		if r != '\uFEFF' {
			br.UnreadRune() // Not a BOM -- put the rune back
		}
		fileContent := ""
		for {
			str, err := br.ReadString('\n')
			fileContent = fileContent + str
			if err == io.EOF {
				break
			}
		}
		log.Println("fileContent:", fileContent)

		if fd.Close() != nil {
			log.Println("file.close err:", err.Error())
		}

		//判断格式
		if !checkData(fileContent) {
			log.Println("check data failed")
			return
		}

		//解析完成后通过http协议发送到服务端
		httpPost(fileContent)

	})
	c.Start()
	select {} //阻塞主线程不退出
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
