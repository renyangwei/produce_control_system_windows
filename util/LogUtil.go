// LogUtil
package util

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

/*
打印日志
*/
func PrintLog(v ...interface{}) {
	if strings.EqualFold("0", ParamDebug()) {
		var fileName = time.Now().Format("2006-01-02") + ".log"
		outfile, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开
		if err != nil {
			fmt.Println(*outfile, "open failed")
			return
		}
		log.SetOutput(outfile)
		log.SetFlags(log.Ldate | log.Ltime)
		log.Println(v)
	}

}
