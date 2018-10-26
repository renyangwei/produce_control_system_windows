// fileUtil
package util

import (
	"io/ioutil"
	"log"
)

func OpenAppConf() string {
	fileContentByte, err := ioutil.ReadFile("app.conf")
	if err != nil {
		PrintLog("read app.conf err:", err.Error())
	}
	fileContent := string(fileContentByte)
	log.Println(fileContent)
	return fileContent
}
