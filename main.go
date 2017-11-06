// PaperManagementClient project main.go
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"PaperManagementClient/sql"
	"PaperManagementClient/util"

	"github.com/robfig/cron"
)

const (
	FILE_PATH string = "path.txt"
)

//文件内容
type DataJson struct {
	Factory string `json:"Factory"`
	Other   string `json:"Other"`
	Time    string `json:"Time"`
	Class   string `json:"Class"`
	Group   string `json:"Group"`
}

//返回数据
type ResponseJson struct {
	Factory string `json:"Factory"`
	Group   string `json:"Group"`
	Class   string `json:"Class"`
	Time    string `json:"Time"`
}

var factoryName string

func main() {

	cronFile()

}

/*
定时读取文件
*/
func cronFile() {
	time_interval := util.Param("time_interval")
	util.PrintLog("time_interval:", time_interval)
	speci := "*/" + time_interval + " * * * * *"
	util.PrintLog("spec:", speci)

	c := cron.New()
	//秒 分 时 日 月 星期
	spec := "*/5 * * * * *" //每五秒一次
	c.AddFunc(spec, func() {
		initFile()
	})
	c.AddFunc(speci, func() {
		sql.Connect()
	})
	c.AddFunc(speci, func() {
		sql.ReadSearchRequest()
	})
	c.Start()
	select {} //阻塞主线程不退出
}

/*
初始化读取文件
*/
func initFile() {
	FILE_NAME_ARRAY := []string{"infor.txt", "infor1.txt", "infor2.txt", "infor3.txt", "infor4.txt", "infor5.txt"}
	FILE_NAME_DATA_ARRAY := []string{"data.txt", "data1.txt", "data2.txt", "data3.txt", "data4.txt", "data5.txt"}
	GROUP_NAME_ARRAY := []string{"一号线", "二号线", "三号线", "四号线", "五号线", "六号线"}

	//读取文件路径
	path, err := os.Open(FILE_PATH)
	if err != nil {
		util.PrintLog("open file path error,", err.Error())
		return
	}
	reader := bufio.NewReader(path)
	r, _, err := reader.ReadRune()
	if err != nil {
		util.PrintLog("read rune err:", err.Error())
		return
	}
	if r != '\uFEFF' {
		reader.UnreadRune() // Not a BOM -- put the rune back
	}
	pathContent, err := reader.ReadString('\n')
	if err != nil {
		util.PrintLog("read string path error,", err.Error())
		return
	}
	util.PrintLog("pathContent:", pathContent)
	pathArray := strings.Split(pathContent, ",")
	for i := 0; i < len(pathArray); i++ {
		pathArr := pathArray[0]
		if i == len(pathArray)-1 {
			pathArr = pathArray[i][0 : len(pathArray[i])-2]
		}
		util.PrintLog("pathArr,", pathArr)
		fileName := pathArr + FILE_NAME_ARRAY[i]
		dataName := pathArr + FILE_NAME_DATA_ARRAY[i]
		readFile(fileName, GROUP_NAME_ARRAY[i], util.GetFactoryUrl())
		readFile(dataName, GROUP_NAME_ARRAY[i], util.GetHistoryUrl())
	}
}

/*
读取文件
fileName	文件名
group		班组
*/
func readFile(fileName string, group string, httpUrl string) {
	//读取文件
	fd, err := os.Open(fileName)
	if err != nil {
		util.PrintLog("open file err:", err.Error())
		return
	}
	br := bufio.NewReader(fd)
	r, _, err := br.ReadRune()
	if err != nil {
		util.PrintLog("read rune err:", err.Error())
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
	util.PrintLog("fileContent:", fileContent)
	if fd.Close() != nil {
		util.PrintLog("file.close err:", err.Error())
		return
	}
	//增加产线
	fileContent = addGroup(fileContent, group)
	//解析完成后通过http协议发送到服务端
	httpPost(fileContent, httpUrl)
}

/*
添加产线
*/
func addGroup(data string, group string) string {
	var fileContentJson DataJson
	err := json.Unmarshal([]byte(data), &fileContentJson)
	if err != nil {
		return err.Error()
	}
	fileContentJson.Group = group
	fileContent, err := json.Marshal(fileContentJson)
	if err != nil {
		return err.Error()
	}
	//保存
	factoryName = fileContentJson.Factory
	return string(fileContent)
}

/*
发送到服务端
*/
func httpPost(data string, httpUrl string) {
	body := bytes.NewBuffer([]byte(data))
	resp, err := http.Post(httpUrl, util.HTTP_APPLICATION, body)
	if err != nil {
		util.PrintLog(err.Error())
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		util.PrintLog(err.Error())
	}
	util.PrintLog(string(respBody))
	parseIsForceRefresh(respBody)
}

/*
增加判断返回内容
*/
func parseIsForceRefresh(response []byte) {
	//读取文件路径
	path, err := os.Open(FILE_PATH)
	if err != nil {
		util.PrintLog("open file path error,", err.Error())
		return
	}
	reader := bufio.NewReader(path)
	r, _, err := reader.ReadRune()
	if err != nil {
		util.PrintLog("read rune err:", err.Error())
		return
	}
	if r != '\uFEFF' {
		reader.UnreadRune() // Not a BOM -- put the rune back
	}
	pathContent, err := reader.ReadString('\n')
	if err != nil {
		util.PrintLog("read string path error,", err.Error())
		return
	}
	util.PrintLog("parseIsForceRefresh pathContent:", pathContent)
	pathArray := strings.Split(pathContent, ",")

	var responseJson ResponseJson
	err = json.Unmarshal(response, &responseJson)
	if err != nil {
		util.PrintLog(err.Error())
		return
	}
	var fileContent string = "{\"Class\":\"" + responseJson.Class + "\", \"Time\": \"" + responseJson.Time + "\"}"
	var fileName string
	//判断文件名
	switch responseJson.Group {
	case "一号线":
		if len(pathArray) > 0 {
			fileName = pathArray[0] + "location.txt"
		}
	case "二号线":
		if len(pathArray) > 1 {
			fileName = pathArray[1] + "location1.txt"
		} else {
			fileName = pathArray[0] + "location1.txt"
		}
	case "三号线":
		if len(pathArray) > 2 {
			fileName = pathArray[2] + "location2.txt"
		} else {
			fileName = pathArray[0] + "location2.txt"
		}
	case "四号线":
		if len(pathArray) > 3 {
			fileName = pathArray[3] + "location3.txt"
		} else {
			fileName = pathArray[0] + "location3.txt"
		}
	case "五号线":
		if len(pathArray) > 4 {
			fileName = pathArray[4] + "location4.txt"
		} else {
			fileName = pathArray[0] + "location4.txt"
		}
	case "六号线":
		if len(pathArray) > 5 {
			fileName = pathArray[5] + "location5.txt"
		} else {
			fileName = pathArray[0] + "location5.txt"
		}
	}
	writeFile(fileName, fileContent)
}

/*
写入文件
*/
func writeFile(fileName string, fileContent string) {
	util.PrintLog("writeFile, fileName:", fileName)
	var d1 = []byte(fileContent)
	err := ioutil.WriteFile(fileName, d1, 0666) //写入文件(字节数组)
	if err != nil {
		util.PrintLog(err.Error())
		return
	}
}
