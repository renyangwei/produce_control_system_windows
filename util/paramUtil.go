// paramUtil
package util

import (
	"strings"
)

var (
	array []string
)

func init() {
	//打开app.conf配置文件
	fileContent := OpenAppConf()
	array = strings.Split(fileContent, "\n")
}

/*
解析参数
*/
func Param(params string) string {
	for _, str := range array {
		if strings.Contains(str, params) && !strings.Contains(str, "#") {
			arr := strings.Split(str, "=")
			return arr[1]
		}
	}
	return ""
}
