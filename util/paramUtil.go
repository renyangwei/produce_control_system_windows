// paramUtil
package util

import (
	"encoding/json"
	"strconv"
	"strings"
)

var (
	//	array []string
	appConf AppConf
)

type AppConf struct {
	Servers      []Db `json:"Server"`
	TimeInterval int  `json:"time_interval"`
	DebugMode    int  `json:"debug"`
	RowsLimit    int  `json:"rows_limit"`
	LocalTest    int  `json:"local_test"`
}

type Db struct {
	Host       string     `json:"host"`
	User       string     `json:"user"`
	Pwd        string     `json:"pwd"`
	Port       int        `json:"port"`
	ServerName string     `json:"server_name"`
	Datas      []Database `json:"database"`
}

type Database struct {
	Name  string `json:"name"`
	Group string `json:"group"`
}

func init() {
	//打开app.conf配置文件
	//fileContent := OpenAppConf()
	//	array = strings.Split(fileContent, "\n")
	//	for _, str := range array {
	//		if strings.Contains(str, "debug") && !strings.Contains(str, "#") {
	//			arr := strings.Split(str, "=")
	//			Debug = Trim(arr[1])
	//		}
	//	}

	fileContent := OpenAppConf()
	//	PrintLog("appconfig fileContent:", fileContent)
	err := json.Unmarshal([]byte(fileContent), &appConf)
	if err != nil {
		PrintLog("unmarshal appconfig error:", err.Error())
		return
	}
}

/*
解析参数
*/
func Param(params string) string {
	//	for _, str := range array {
	//		if strings.Contains(str, params) && !strings.Contains(str, "#") {
	//			arr := strings.Split(str, "=")
	//			return Trim(arr[1])
	//		}
	//	}
	if params == "local_test" {
		return ParamLocalTest()
	}
	return ""
}

/*
获取数据库参数
*/
func ParamServers() []Db {
	return appConf.Servers
}

/*
获取时间间隔
*/
func ParamTimeInterval() string {
	return strconv.Itoa(appConf.TimeInterval)
}

/*
获得debug参数
*/
func ParamDebug() string {
	return strconv.Itoa(appConf.DebugMode)
}

/*
获得查询条数
*/
func ParamRowsLimit() string {
	return strconv.Itoa(appConf.RowsLimit)
}

/*
获得是否本地测试
*/
func ParamLocalTest() string {
	return strconv.Itoa(appConf.LocalTest)
}

/*
去掉空格和回车符号
*/
func Trim(old string) string {
	newStr := strings.Replace(old, "\r", "", -1)
	newStr = strings.Replace(newStr, "\n", "", -1)
	newStr = strings.Replace(newStr, " ", "", -1)
	return newStr
}
