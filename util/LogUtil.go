// LogUtil
package util

import (
	"log"
)

//var debug string

//func init() {
//	debug = Param("debug")
//	log.Println("debug:", debug)
//}

/*
打印日志
*/
func PrintLog(v ...interface{}) {
	debug := Param("debug")
	//	log.Println("debug:", debug)
	if debug == "0" {
		log.Println(v)
	}
}
