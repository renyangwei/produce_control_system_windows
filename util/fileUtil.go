// fileUtil
package util

import (
	"io/ioutil"
)

func OpenAppConf() string {
	fileContentByte, err := ioutil.ReadFile("app.conf")
	if err != nil {
		PrintLog("read app.conf err:", err.Error())
	}
	fileContent := string(fileContentByte)
	PrintLog(fileContent)
	return fileContent
}
