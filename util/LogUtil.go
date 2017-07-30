// LogUtil
package util

import (
	"log"
	"strings"
)

/*
打印日志
*/
func PrintLog(v ...interface{}) {
	//	log.Println("debug:", Debug)
	if strings.EqualFold(Debug, "0") {
		log.Println(v)
	}
}
