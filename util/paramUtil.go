// paramUtil
package util

import (
	"strings"
)

var (
	array []string
	Debug string
)

func init() {
	//打开app.conf配置文件
	fileContent := OpenAppConf()
	array = strings.Split(fileContent, "\n")
	for _, str := range array {
		if strings.Contains(str, "debug") && !strings.Contains(str, "#") {
			arr := strings.Split(str, "=")
			Debug = Trim(arr[1])
		}
	}
}

/*
解析参数
*/
func Param(params string) string {
	for _, str := range array {
		if strings.Contains(str, params) && !strings.Contains(str, "#") {
			arr := strings.Split(str, "=")
			return Trim(arr[1])
		}
	}
	return ""
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
