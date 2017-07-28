// sqlmanage
package sql

import (
	"PaperManagementClient/util"
	"strconv"
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
	ConnectSqlServer(host_, user_, pwd_, database_, server_name_, iport, rows_limit_)
}
