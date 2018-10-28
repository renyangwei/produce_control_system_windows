// sqlmanage
package sql

import (
	"PaperManagementClient/util"
)

var (
	rows_limit_ string
	dbs         []util.Db
)

func init() {
	rows_limit_ = util.ParamRowsLimit()
	dbs = util.ParamServers()
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
连接数据库,读取订单和完工资料
*/
func Connect() {
	util.PrintLog("dbs", dbs)
	for _, db := range dbs {
		if db.Host != "" {
			var orderSqlSyn = "select top " + rows_limit_ + " a.scxh,a.mxbh,a.khjc,zbdh=left(rtrim(a.zbdh)+'--------------',7*c.zlbhcd),a.klzhdh,a.zd,a.zbcd,pcsl=a.ddsl-isnull(a.tlsl,0),a.ddsm,a.zt,a.ks, a.sm2,a.zbcd2,xbmm=round((a.zd-a.jbkd)*10/c.convertvalue,0),scbh=isnull(a.scbh,''),ms=((a.ddsl-a.tlsl)*a.zbcd)/c.convertvalue,a.finishtime from xddmx a,xtsz c where a.zt in (1,2) and a.ddsl-isnull(a.tlsl,0)>0 and isnull(a.cczt,0)<9 order by a.zt desc,a.scxh,a.zdxh,a.zbxh,a.zd desc,a.zbdh,a.khbh,a.zbcd desc"
			var finishInfoSqlSyn = "select top " + rows_limit_ + " a.mxbh, a.khjc,a.pcsl,a.hgpsl,a.blpsl,a.zd,a.zbmc,a.zbcd,xdzd=a.zbkd/a.ks,xbmm=round((a.zd-a.zbkd)*10/b.convertvalue,0),a.klzhdh,a.ks,a.stoptime,a.stopspec,a.bzbh,a.starttime,a.finishtime,ys=case when convert(char(19),a.starttime,21)<convert(char(19),a.finishtime,21) then  datediff(s,a.starttime,a.finishtime)  else  0  end,  a.zbcd2,shl=case when (a.hgpsl+a.blpsl)>0 then str(round(a.blpsl*100.0/(a.hgpsl+a.blpsl),2),4,2)+'%' else '0%' end,js=case when convert(char(19),a.starttime,21)<convert(char(19),a.finishtime,21) then round(60*a.zbcd*(a.hgpsl+a.blpsl)/(100*b.convertvalue)/datediff(s,a.starttime,a.finishtime),0) else 0 End,ms=round(a.zbcd*a.hgpsl/(100*b.convertvalue),0) from finish a,xtsz b where convert(char(19),a.starttime,21)>='2000-01-01' and convert(char(19),a.finishtime,21)<= '2050-08-19' and isnull(a.pcsl,0)>0 and a.khjc<>'' order by a.finishtime desc,a.starttime desc,a.scxh desc"
			var histroySqlSyn = "select top " + rows_limit_ + " qsrq, jzrq, tjsj, pjjs, pjzd, dds, hlcs, zms, zhgms, zmj, zhgmj, tjcs, zzl, hgzl, zxbmj, zxbzl, xbbl, bzbh, rq from schzb order by rq desc"
			util.PrintLog("db.Datas", db.Datas)
			for _, data := range db.Datas {
				ConnectSqlServer(db.Host, db.User, db.Pwd, data.Name, db.ServerName, db.Port, rows_limit_, data.Group, orderSqlSyn, finishInfoSqlSyn, histroySqlSyn)
			}
		}
	}
}

/*
搜索参数
*/
func ReadSearchRequest() {
	for _, db := range dbs {
		if db.Host != "" {
			for _, data := range db.Datas {
				SearchSqlServer(db.Host, db.User, db.Pwd, data.Name, db.ServerName, db.Port, rows_limit_, data.Group)
			}
		}
	}
}
