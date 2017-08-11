// sqlmanage
package sql

import (
	"PaperManagementClient/util"
	"strconv"
	"strings"
)

var (
	host_        string
	user_        string
	pwd_         string
	port_        string
	database_    string
	server_name_ string
	rows_limit_  string
	Debug        string
)

func init() {
	host_ = util.Param("host")
	user_ = util.Param("user")
	pwd_ = util.Param("pwd")
	port_ = util.Param("port")
	database_ = util.Param("database")
	server_name_ = util.Param("server_name")
	rows_limit_ = util.Param("rows_limit")
	Debug = util.Param("debug")
}

func Connect() {
	iport, err := strconv.Atoi(port_)
	if err != nil {
		util.PrintLog("params port err:", err.Error())
		return
	}
	//读取不同的数据库
	dbs := strings.Split(database_, ",")
	util.PrintLog("dbs:", dbs)
	for _, db := range dbs {
		dbg := strings.Split(db, ":")
		if !strings.EqualFold(dbg[0], "") && !strings.EqualFold(dbg[1], "") {
			ConnectSqlServer(host_, user_, pwd_, dbg[0], server_name_, iport, rows_limit_, dbg[1])
		}
	}
}
