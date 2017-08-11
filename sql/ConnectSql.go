// ConnectSql
package sql

import (
	"PaperManagementClient/util"
	"database/sql"
	"encoding/json"
	"strconv"
	"time"

	_ "github.com/mattn/go-adodb"
)

var (
	local bool = false

	host        string
	user        string
	pwd         string
	port        int
	databases   string
	server_name string
	rows_limit  string
	group       string

	cname string
)

type Mssql struct {
	*sql.DB
	dataSource string
	database   string
	windows    bool
	sa         *SA
}

type SA struct {
	user   string
	passwd string
	port   int
}

func NewMssql() *Mssql {
	mssql := new(Mssql)
	dataS := host + "\\" + server_name
	mssql = &Mssql{
		// 如果数据库是默认实例（MSSQLSERVER）则直接使用IP，命名实例需要指明。
		// dataSource: "192.168.15.128\\MSSQLSERVER",
		dataSource: dataS,
		database:   databases,
		// windows: true 为windows身份验证，false 必须设置sa账号和密码
		windows: local,
		sa: &SA{
			user:   user,
			passwd: pwd,
			port:   port,
		},
	}
	return mssql
}

func (m *Mssql) Open() error {
	config := "Provider=SQLOLEDB;Initial Catalog=" + m.database + ";Data Source=" + m.dataSource

	if m.windows {
		config = config + ";Integrated Security=SSPI"
	} else {
		// sql 2000的端口写法和sql 2005以上的有所不同，在Data Source 后以逗号隔开。
		config = config + "," + strconv.Itoa(m.sa.port) + ";user id=" + m.sa.user + ";password=" + m.sa.passwd
	}
	util.PrintLog(config)
	var err error
	m.DB, err = sql.Open("adodb", config)
	return err
}

/*
查询公司名
*/
func (m *Mssql) SelectCompany() {
	rows, err := m.Query("select yhmc from xtsz")
	if err != nil {
		util.PrintLog("select cname err:", err.Error())
		return
	}
	for rows.Next() {
		rows.Scan(&cname)
	}
	util.PrintLog("cname:", cname)
}

/*
读取并发送订单资料
*/
func (m *Mssql) SelectOrder() {
	rows, err := m.Query("select top " + rows_limit + " a.scxh,a.mxbh,a.khjc,zbdh=left(rtrim(a.zbdh)+'--------------',7*c.zlbhcd),a.klzhdh,a.zd,a.zbcd,pcsl=a.ddsl-isnull(a.tlsl,0),a.ddsm,a.zt,a.ks, a.sm2,a.zbcd2,xbmm=round((a.zd-a.jbkd)*10/c.convertvalue,0),a.scbh,ms=round((a.ddsl-a.tlsl)*a.zbcd/(c.convertvalue*100),0),a.finishtime from xddmx a,xtsz c where a.zt in (1,2) and a.ddsl-isnull(a.tlsl,0)>0 and isnull(a.cczt,0)<9 order by a.zt desc,a.scxh,a.zdxh,a.zbxh,a.zd desc,a.zbdh,a.khbh,a.zbcd desc")
	if err != nil {
		util.PrintLog("select query err: %s\n", err)
		return
	}
	if cname == "" {
		util.PrintLog("cname is empty")
		return
	}
	var normalDatas []util.NormarData
	for rows.Next() {
		var (
			mxbh        string    //订单号
			khjc        string    //客户简称
			zbdh        string    //材质
			klzhdh      string    //楞别
			xdzd        string    //纸度
			pscl        string    //排产数量
			zbcd        string    //切长
			ks          string    //剖
			finish_time time.Time //预计完工时间
			ddms        string    //留言
		)
		rows.Scan(&mxbh, &khjc, &zbdh, &klzhdh, &xdzd, &pscl, &zbcd, &ks, &finish_time, &ddms)
		var order util.Order
		order.Mxbh = util.Trim(mxbh)
		order.Khjc = util.Trim(khjc)
		order.Zbdh = util.Trim(zbdh)
		order.Klzhdh = util.Trim(klzhdh)
		order.Xdzd = util.Trim(xdzd)
		order.Pscl = util.Trim(pscl)
		order.Zbcd = util.Trim(zbcd)
		order.Ks = util.Trim(ks)
		order.Finish_time = finish_time
		order.Ddms = util.Trim(ddms)
		orderJson, err := json.Marshal(order)
		if err != nil {
			util.PrintLog(err.Error())
			return
		}
		var normalData util.NormarData
		normalData.Cname = util.Trim(cname)
		normalData.Data = string(orderJson)
		normalData.Group = group
		normalDatas = append(normalDatas, normalData)
	}
	datasJson, err := json.Marshal(normalDatas)
	if err != nil {
		util.PrintLog(err.Error())
		return
	}
	util.PostOrder(string(datasJson))
}

/*
读取并发送完工资料
*/
func (m *Mssql) selectFinishInfo() {
	rows, err := m.Query("select top " + rows_limit + " mxbh,khjc,zbbh,zbkd,hgpsl,blpsl,pcsl,zbcd,starttime,finishtime from finish where hgpsl>0 or blpsl>0 order by finishtime desc")
	if err != nil {
		util.PrintLog("select query err:", err)
		return
	}
	if cname == "" {
		util.PrintLog("cname is empty")
		return
	}
	var normalDatas []util.NormarData
	for rows.Next() {
		var (
			mxbh        string    //订单号
			khjc        string    //客户简称
			zbdh        string    //材质
			zbkd        string    //纸板宽
			hgpsl       string    //合格数
			blpsl       string    //不良数
			pcsl        string    //排产数
			zbcd        string    //切长
			start_time  time.Time //开始时间
			finish_time time.Time //完工时间
		)
		rows.Scan(&mxbh, &khjc, &zbdh, &zbkd, &hgpsl, &blpsl, &pcsl, &zbcd, &start_time, &finish_time)
		var finishInfo util.FinishInfo
		finishInfo.Mxbh = util.Trim(mxbh)
		finishInfo.Khjc = util.Trim(khjc)
		finishInfo.Zbdh = util.Trim(zbdh)
		finishInfo.Hgpsl = util.Trim(hgpsl)
		finishInfo.Blpsl = util.Trim(blpsl)
		finishInfo.Pcsl = util.Trim(pcsl)
		finishInfo.Zbcd = util.Trim(zbcd)
		finishInfoJson, err := json.Marshal(finishInfo)
		if err != nil {
			util.PrintLog("marshal json err:", err.Error())
			return
		}
		var normalData util.NormarData
		normalData.Cname = util.Trim(cname)
		normalData.Data = string(finishInfoJson)
		normalData.StartTime = start_time
		normalData.FinishTime = finish_time
		normalData.Group = group
		normalDatas = append(normalDatas, normalData)
	}
	datasJson, err := json.Marshal(normalDatas)
	if err != nil {
		util.PrintLog(err.Error())
		return
	}
	util.PostFinihInfo(string(datasJson))
}

/*
连接sqlserver
*/
func ConnectSqlServer(_host, _user, _pwd, _database, _server_name string, _port int, _rows_limit, _group string) {

	host = _host
	user = _user
	pwd = _pwd
	port = _port
	databases = _database
	server_name = _server_name
	rows_limit = _rows_limit
	group = _group

	mssql := NewMssql()
	err := mssql.Open()
	if err != nil {
		util.PrintLog(err)
	}

	mssql.SelectCompany()
	mssql.SelectOrder()
	mssql.selectFinishInfo()
}
