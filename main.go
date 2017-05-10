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
	//	HTTP_HOST string = "http://180.76.163.58:8081/factory"
	HTTP_HOST        string = "http://localhost:8081/factory"
	HTTP_APPLICATION string = "application/json;charset=utf-8"
	FILE_NAME        string = "infor.txt"
	GROUP_NAME       string = "一号线"
	FILE_NAME_1      string = "infor1.txt"
	GROUP_NAME_1     string = "二号线"
	FILE_NAME_2      string = "infor2.txt"
	GROUP_NAME_2     string = "三号线"
)

type RealTimeDataJson struct {
	Factory string `json:"Factory"`
	Other   string `json:"Other"`
	Group   string `json:"Group"`
}

func main() {
	c := cron.New()
	//秒 分 时 日 月 星期
	//	spec := "0 */1 * * * *" //每分钟一次
	spec := "*/5 * * * * *" //每五秒一次
	c.AddFunc(spec, func() {
		readFile(FILE_NAME, GROUP_NAME)
		readFile(FILE_NAME_1, GROUP_NAME_1)
		readFile(FILE_NAME_2, GROUP_NAME_2)
	})
	c.Start()
	select {} //阻塞主线程不退出
}

/*
读取文件
fileName	文件名
group		班组
*/
func readFile(fileName string, group string) {
	//读取文件
	fd, err := os.Open(fileName)
	if err != nil {
		log.Println("open file err:", err.Error())
		return
	}
	br := bufio.NewReader(fd)
	r, _, err := br.ReadRune()
	if err != nil {
		log.Println("read rune err:", err.Error())
		return
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
		return
	}

	//增加产线
	fileContent = addGroup(fileContent, group)

	//解析完成后通过http协议发送到服务端
	httpPost(fileContent)
}

/*
添加产线
*/
func addGroup(data string, group string) string {
	var fileContentJson RealTimeDataJson
	err := json.Unmarshal([]byte(data), &fileContentJson)
	if err != nil {
		log.Println("json data invalid")
		return err.Error()
	}
	fileContentJson.Group = group
	fileContent, err := json.Marshal(fileContentJson)
	if err != nil {
		log.Println("json to string failed")
		return err.Error()
	}
	return string(fileContent)
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
		log.Println(err)
	}

	log.Println(string(respBody))

}
