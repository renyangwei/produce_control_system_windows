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

type SearchRequestStruct struct {
	Type       string //搜索类型:order,finish_info
	Cname      string //公司名
	Group      string //产线
	Data       string //搜索内容
	StartTime  string
	FinishTime string
}

/*
连接数据库
*/
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
			var orderSqlSyn = "select top " + rows_limit_ + " a.scxh,a.mxbh,a.khjc,zbdh=left(rtrim(a.zbdh)+'--------------',7*c.zlbhcd),a.klzhdh,a.zd,a.zbcd,pcsl=a.ddsl-isnull(a.tlsl,0),a.ddsm,a.zt,a.ks, a.sm2,a.zbcd2,xbmm=round((a.zd-a.jbkd)*10/c.convertvalue,0),scbh=isnull(a.scbh,''),ms=((a.ddsl-a.tlsl)*a.zbcd)/c.convertvalue,a.finishtime from xddmx a,xtsz c where a.zt in (1,2) and a.ddsl-isnull(a.tlsl,0)>0 and isnull(a.cczt,0)<9 order by a.zt desc,a.scxh,a.zdxh,a.zbxh,a.zd desc,a.zbdh,a.khbh,a.zbcd desc"
			var finishInfoSqlSyn = "select top " + rows_limit_ + " a.mxbh, a.khjc,a.pcsl,a.hgpsl,a.blpsl,a.zd,a.zbmc,a.zbcd,xdzd=a.zbkd/a.ks,xbmm=round((a.zd-a.zbkd)*10/b.convertvalue,0),a.klzhdh,a.ks,a.stoptime,a.stopspec,a.bzbh,a.starttime,a.finishtime,ys=case when convert(char(19),a.starttime,21)<convert(char(19),a.finishtime,21) then  datediff(s,a.starttime,a.finishtime)  else  0  end,  a.zbcd2,shl=case when (a.hgpsl+a.blpsl)>0 then str(round(a.blpsl*100.0/(a.hgpsl+a.blpsl),2),4,2)+'%' else '0%' end,js=case when convert(char(19),a.starttime,21)<convert(char(19),a.finishtime,21) then round(60*a.zbcd*(a.hgpsl+a.blpsl)/(100*b.convertvalue)/datediff(s,a.starttime,a.finishtime),0) else 0 End,ms=round(a.zbcd*a.hgpsl/(100*b.convertvalue),0) from finish a,xtsz b where convert(char(19),a.starttime,21)>='2000-01-01' and convert(char(19),a.finishtime,21)<= '2050-08-19' and isnull(a.pcsl,0)>0 and a.khjc<>'' order by a.finishtime desc,a.starttime desc,a.scxh desc"
			ConnectSqlServer(host_, user_, pwd_, dbg[0], server_name_, iport, rows_limit_, dbg[1], orderSqlSyn, finishInfoSqlSyn)
		}
	}
}

/*
读取搜索参数
*/
func ReadSearchRequest() {
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
			SearchSqlServer(host_, user_, pwd_, dbg[0], server_name_, iport, rows_limit_, dbg[1])
		}
	}
}
