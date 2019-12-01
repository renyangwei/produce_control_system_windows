// ConnectSql
package sql

import (
	"database/sql"
	"encoding/json"
	"produce_control_system_windows/util"
	"strconv"

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
	//	group       string

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

var mssql *Mssql

func init() {
	util.PrintLog("connectSql init")
	mssql = NewMssql()
	err := mssql.Open()
	if err != nil {
		util.PrintLog(err)
		return
	}
}

//返回数据
type HistoryResponseJson struct {
	Factory string `json:"Factory"`
	Group   string `json:"Group"`
	Class   string `json:"Class"`
	Time    string `json:"Time"`
}

func NewMssql() *Mssql {
	sql := new(Mssql)
	dataS := host + "\\" + server_name
	sql = &Mssql{
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
	return sql
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
func (m *Mssql) selectCompany() {
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
func (m *Mssql) selectOrder(sqlSyn, gup string) string {
	util.PrintLog("order sql:" + sqlSyn)
	rows, err := m.Query(sqlSyn)
	if err != nil {
		util.PrintLog("select query err: %s\n", err)
		return ""
	}
	if cname == "" {
		util.PrintLog("cname is empty")
		return ""
	}
	var normalDatas []util.NormarData
	for rows.Next() {
		var (
			scxh        string //序号
			mxbh        string //订单号
			khjc        string //客户简称
			zbdh        string //材质
			klzhdh      string //楞别
			zd          string //纸度
			zbcd        string //切长
			pscl        string //排产数量
			ddms        string //留言
			zt          string
			ks          string //剖
			sm2         string
			zbcd2       string
			xbmm        string
			scbh        string
			ms          float64
			finish_time string //预计完工时间
		)
		rows.Scan(&scxh, &mxbh, &khjc, &zbdh, &klzhdh, &zd, &zbcd, &pscl, &ddms, &zt, &ks, &sm2, &zbcd2, &xbmm, &scbh, &ms, &finish_time)
		var order util.Order
		order.Scxh = util.Trim(scxh)
		order.Mxbh = util.Trim(mxbh)
		order.Khjc = util.Trim(khjc)
		order.Zbdh = util.Trim(zbdh)
		order.Klzhdh = util.Trim(klzhdh)
		order.Zd = util.Trim(zd)
		order.Zbcd = util.Trim(zbcd)
		order.Pscl = util.Trim(pscl)
		order.Ddms = util.Trim(ddms)
		order.Zt = util.Trim(zt)
		order.Ks = util.Trim(ks)
		order.Sm2 = util.Trim(sm2)
		order.Zbcd2 = util.Trim(zbcd2)
		order.Xbmm = util.Trim(xbmm)
		order.Scbh = util.Trim(scbh)
		order.Ms = ms
		order.Finish_time = finish_time
		orderJson, err := json.Marshal(order)
		if err != nil {
			util.PrintLog(err.Error())
			return ""
		}
		var normalData util.NormarData
		normalData.Cname = util.Trim(cname)
		normalData.Data = string(orderJson)
		normalData.Group = gup
		normalDatas = append(normalDatas, normalData)
	}
	datasJson, err := json.Marshal(normalDatas)
	if err != nil {
		util.PrintLog(err.Error())
		return ""
	}
	return string(datasJson)
}

/*
读取并发送完工资料
*/
func (m *Mssql) selectFinishInfo(sqlSyn, gup string) string {
	util.PrintLog("finish info sqlSyn:", sqlSyn)
	rows, err := m.Query(sqlSyn)
	if err != nil {
		util.PrintLog("select query err:", err)
		return ""
	}
	if cname == "" {
		util.PrintLog("cname is empty")
		return ""
	}
	var normalDatas []util.NormarData
	for rows.Next() {
		var (
			mxbh        string //订单编号
			khjc        string //客户简称
			pcsl        string //排产数
			hgpsl       string //合格数
			blpsl       string //不良数
			zd          string
			zbmc        string
			zbcd        string //切长
			xdzd        string
			xbmm        string
			klzhdh      string
			ks          string
			stop_time   string
			stop_spec   string
			bzbh        string
			start_time  string //开始时间
			finish_time string //完工时间
			ys          string
			zbcd2       string
			shl         string
			js          string
			ms          string
		)
		rows.Scan(&mxbh, &khjc, &pcsl, &hgpsl, &blpsl, &zd, &zbmc, &zbcd, &xdzd, &xbmm, &klzhdh, &ks, &stop_time, &stop_spec, &bzbh, &start_time, &finish_time, &ys, &zbcd2, &shl, &js, &ms)
		var finishInfo util.FinishInfo
		finishInfo.Mxbh = util.Trim(mxbh)
		finishInfo.Khjc = util.Trim(khjc)
		finishInfo.Pcsl = util.Trim(pcsl)
		finishInfo.Hgpsl = util.Trim(hgpsl)
		finishInfo.Blpsl = util.Trim(blpsl)
		finishInfo.Zd = util.Trim(zd)
		finishInfo.Zbmc = util.Trim(zbmc)
		finishInfo.Xbmm = util.Trim(xbmm)
		finishInfo.Klzhdh = util.Trim(klzhdh)
		finishInfo.Ks = util.Trim(ks)
		finishInfo.StopTime = util.Trim(stop_time)
		finishInfo.StopSpec = util.Trim(stop_spec)
		finishInfo.Bzbh = util.Trim(bzbh)
		finishInfo.Zbcd2 = util.Trim(zbcd2)
		finishInfo.Ys = util.Trim(ys)
		finishInfo.Shl = util.Trim(shl)
		finishInfo.Js = util.Trim(js)
		finishInfo.Ms = util.Trim(ms)
		finishInfoJson, err := json.Marshal(finishInfo)
		if err != nil {
			util.PrintLog("marshal json err:", err.Error())
			return ""
		}
		var normalData util.NormarData
		normalData.Cname = util.Trim(cname)
		normalData.Data = string(finishInfoJson)
		normalData.StartTime = start_time
		normalData.FinishTime = finish_time
		normalData.Group = gup
		normalDatas = append(normalDatas, normalData)
	}
	datasJson, err := json.Marshal(normalDatas)
	if err != nil {
		util.PrintLog(err.Error())
		return ""
	}
	return string(datasJson)
}

/*
连接sqlserver
*/
func ConnectSqlServer(_host, _user, _pwd, _database, _server_name string, _port int, _rows_limit, _group, orderSqlSyn, finishInfoSqlSyn, historySqlSyn string) {

	host = _host
	user = _user
	pwd = _pwd
	port = _port
	databases = _database
	server_name = _server_name
	rows_limit = _rows_limit

	// mssql := NewMssql()
	// err := mssql.Open()
	// if err != nil {
	// 	util.PrintLog(err)
	// 	return
	// }
	//连接上数据库再赋值
	//group = _group

	mssql.selectCompany()
	orderData := mssql.selectOrder(orderSqlSyn, _group)
	if orderData != "" {
		util.PostOrder(orderData)
	}
	finishInfoData := mssql.selectFinishInfo(finishInfoSqlSyn, _group)
	if finishInfoData != "" {
		util.PostFinihInfo(finishInfoData)
	}
	mssql.queryHistoryData(historySqlSyn, _group)
}

/*
到数据库搜索
*/
func SearchSqlServer(_host, _user, _pwd, _database, _server_name string, _port int, _rows_limit, _group string) {
	host = _host
	user = _user
	pwd = _pwd
	port = _port
	databases = _database
	server_name = _server_name
	rows_limit = _rows_limit

	// mssql := NewMssql()
	// if cname == "" {
	// 	err := mssql.Open()
	// 	if err != nil {
	// 		util.PrintLog(err)
	// 		return
	// 	}
	mssql.selectCompany()
	// }
	if cname == "" {
		util.PrintLog("cname is empty")
		return
	}
	//连接上再赋值
	//	group = _group

	searchRequest := util.GetSearchRequest(util.Trim(cname))
	//解析数据
	var searchRequestStruct SearchRequestStruct
	err := json.Unmarshal([]byte(searchRequest), &searchRequestStruct)
	if err != nil {
		util.PrintLog("ReadSearchRequest, unmarshal search request err:", err)
		return
	}
	if searchRequestStruct.Group != _group {
		util.PrintLog("searchRequestStruct.Group:", searchRequestStruct.Group, ",group:", _group)
		return
	}
	// err = mssql.Open()
	// if err != nil {
	// 	util.PrintLog(err)
	// 	return
	// }
	//根据type到数据里搜索
	if searchRequestStruct.Type == "order" {
		//搜索订单
		var sqlSyn string
		if searchRequestStruct.Data == "" {
			//如果没有输入，那么查找所有的数据
			sqlSyn = "select a.scxh,a.mxbh,a.khjc,zbdh=left(rtrim(a.zbdh)+'--------------',7*c.zlbhcd),a.klzhdh,a.zd,a.zbcd,pcsl=a.ddsl-isnull(a.tlsl,0),a.ddsm,a.zt,a.ks, a.sm2,a.zbcd2,xbmm=round((a.zd-a.jbkd)*10/c.convertvalue,0),scbh=isnull(a.scbh,''),ms=((a.ddsl-a.tlsl)*a.zbcd)/(c.convertvalue),a.finishtime from xddmx a,xtsz c where a.zt in (1,2) and a.ddsl-isnull(a.tlsl,0)>0 and isnull(a.cczt,0)<9 order by a.zt desc,a.scxh,a.zdxh,a.zbxh,a.zd desc,a.zbdh,a.khbh,a.zbcd desc"
		} else {
			sqlSyn = "select a.scxh,a.mxbh,a.khjc,zbdh=left(rtrim(a.zbdh)+'--------------',7*c.zlbhcd),a.klzhdh,a.zd,a.zbcd,pcsl=a.ddsl-isnull(a.tlsl,0),a.ddsm,a.zt,a.ks, a.sm2,a.zbcd2,xbmm=round((a.zd-a.jbkd)*10/c.convertvalue,0),scbh=isnull(a.scbh,''),ms=((a.ddsl-a.tlsl)*a.zbcd)/(c.convertvalue),a.finishtime from xddmx a,xtsz c where a.zt in (1,2) and a.ddsl-isnull(a.tlsl,0)>0 and isnull(a.cczt,0)<9 and (upper(mxbh) like '%" + searchRequestStruct.Data + "%' or lower(mxbh) like '%" + searchRequestStruct.Data + "%' or upper(khjc) like '%" + searchRequestStruct.Data + "%' or lower(khjc) like '%" + searchRequestStruct.Data + "%' or zbcd2 like '%" + searchRequestStruct.Data + "%') order by a.zt desc,a.scxh,a.zdxh,a.zbxh,a.zd desc,a.zbdh,a.khbh,a.zbcd desc"
		}
		searchOrderData := mssql.selectOrder(sqlSyn, _group)
		if searchOrderData != "" {
			//发送到服务器
			util.PrintLog("searchOrderData:", searchOrderData)
			util.PostSearchResult(searchOrderData)
		}
	} else if searchRequestStruct.Type == "finish_info" {
		//搜索完工资料
		var finishInfoSqlSyn string
		if searchRequestStruct.Data == "" {
			//如果没有输入，那么查找所有的数据
			finishInfoSqlSyn = "select a.mxbh, a.khjc,a.pcsl,a.hgpsl,a.blpsl,a.zd,a.zbmc,a.zbcd,xdzd=a.zbkd/a.ks,xbmm=round((a.zd-a.zbkd)*10/b.convertvalue,0),a.klzhdh,a.ks,a.stoptime,a.stopspec,a.bzbh,a.starttime,a.finishtime,ys=case when convert(char(19),a.starttime,21)<convert(char(19),a.finishtime,21) then  datediff(s,a.starttime,a.finishtime)  else  0  end,  a.zbcd2,shl=case when (a.hgpsl+a.blpsl)>0 then str(round(a.blpsl*100.0/(a.hgpsl+a.blpsl),2),4,2)+'%' else '0%' end,js=case when convert(char(19),a.starttime,21)<convert(char(19),a.finishtime,21) then round(60*a.zbcd*(a.hgpsl+a.blpsl)/(100*b.convertvalue)/datediff(s,a.starttime,a.finishtime),0) else 0 End,ms=round(a.zbcd*a.hgpsl/(100*b.convertvalue),0) from finish a,xtsz b where convert(char(19),a.starttime,21)>='" + searchRequestStruct.StartTime + "' and convert(char(19),a.finishtime,21)<= '" + searchRequestStruct.FinishTime + "' and isnull(a.pcsl,0)>0 and a.khjc<>'' order by a.finishtime desc,a.starttime desc,a.scxh desc"
		} else {
			finishInfoSqlSyn = "select a.mxbh, a.khjc,a.pcsl,a.hgpsl,a.blpsl,a.zd,a.zbmc,a.zbcd,xdzd=a.zbkd/a.ks,xbmm=round((a.zd-a.zbkd)*10/b.convertvalue,0),a.klzhdh,a.ks,a.stoptime,a.stopspec,a.bzbh,a.starttime,a.finishtime,ys=case when convert(char(19),a.starttime,21)<convert(char(19),a.finishtime,21) then  datediff(s,a.starttime,a.finishtime)  else  0  end,  a.zbcd2,shl=case when (a.hgpsl+a.blpsl)>0 then str(round(a.blpsl*100.0/(a.hgpsl+a.blpsl),2),4,2)+'%' else '0%' end,js=case when convert(char(19),a.starttime,21)<convert(char(19),a.finishtime,21) then round(60*a.zbcd*(a.hgpsl+a.blpsl)/(100*b.convertvalue)/datediff(s,a.starttime,a.finishtime),0) else 0 End,ms=round(a.zbcd*a.hgpsl/(100*b.convertvalue),0) from finish a,xtsz b where convert(char(19),a.starttime,21)>='" + searchRequestStruct.StartTime + "' and convert(char(19),a.finishtime,21)<= '" + searchRequestStruct.FinishTime + "' and (upper(mxbh) like '%" + searchRequestStruct.Data + "%' or lower(mxbh) like '%" + searchRequestStruct.Data + "%' or upper(khjc) like '%" + searchRequestStruct.Data + "%' or lower(khjc) like '%" + searchRequestStruct.Data + "%' or zbcd2 like '%" + searchRequestStruct.Data + "%') and isnull(a.pcsl,0)>0 and a.khjc<>'' order by a.finishtime desc,a.starttime desc,a.scxh desc"
		}
		var finishInfoData = mssql.selectFinishInfo(finishInfoSqlSyn, _group)
		if finishInfoData != "" {
			util.PrintLog("finishInfoData", finishInfoData)
			//发送到服务器
			util.PostSearchResult(finishInfoData)
		}
	}
}

/*
查询历史订单
*/
func (m *Mssql) queryHistoryData(sqlSyn, gup string) {
	util.PrintLog("history sqlSyn:", sqlSyn)
	rows, err := m.Query(sqlSyn)
	if err != nil {
		util.PrintLog("select query err:", err)
		return
	}
	if cname == "" {
		util.PrintLog("cname is empty")
		return
	}
	for rows.Next() {
		var (
			Qsrq  string //开工时间
			Jzrq  string //完成时间
			Tjsj  string //停机时间
			Pjjs  string //平均车速
			Pjzd  string //平均门幅
			Dds   string //订单笔数
			Hlcs  string //换楞次数
			Zms   string //累计米数
			Zhgms string //合格米数
			Zmj   string //累计面积
			Zhgmj string //合格面积
			Tjcs  string //停机次数
			Zzl   string //累计重量
			Hgzl  string //合格重量
			Zxbmj string //累计修边(m2)
			Zxbzl string //累计修边(kg)
			Xbbl  string //修边比例
			Bzbh  string //班组
			Rq    string //日期
		)
		rows.Scan(&Qsrq, &Jzrq, &Tjsj, &Pjjs, &Pjzd, &Dds, &Hlcs, &Zms, &Zhgms, &Zmj, &Zhgmj, &Tjcs, &Zzl, &Hgzl, &Zxbmj, &Zxbzl, &Xbbl, &Bzbh, &Rq)
		var histroyOther util.OtherStruct
		histroyOther.Qsrq = Qsrq
		histroyOther.Jzrq = Jzrq
		histroyOther.Tjsj = Tjsj
		histroyOther.Pjjs = Pjjs
		histroyOther.Dds = Dds
		histroyOther.Hlcs = Hlcs
		histroyOther.Zms = Zms
		histroyOther.Zhgms = Zhgms
		histroyOther.Zmj = Zmj
		histroyOther.Zhgmj = Zhgmj
		histroyOther.Tjcs = Tjcs
		histroyOther.Zzl = Zzl
		histroyOther.Hgzl = Hgzl
		histroyOther.Zxbmj = Zxbmj
		histroyOther.Zxbzl = Zxbzl
		histroyOther.Xbbl = Xbbl
		histroyOther.Pjzd = Pjzd
		util.PrintLog("query histroy:", histroyOther)
		//拼接成发送的数据
		var hisStruct util.HistroyStruct
		hisStruct.Factory = cname
		hisStruct.Class = Bzbh
		hisStruct.Group = gup
		hisStruct.Time = Rq
		hisStruct.Other = getHistoryOther(histroyOther)
		//发送到服务端
		hisJson, err := json.Marshal(hisStruct)
		if err != nil {
			util.PrintLog("marshl json err:", err.Error())
			return
		}
		var resp = util.PostHistory(string(hisJson))
		mashalPostHistory(resp)
	}
}

/*
解析发送历史数据的返回内容
*/
func mashalPostHistory(resp string) {
	var responseJson HistoryResponseJson
	err := json.Unmarshal([]byte(resp), &responseJson)
	if err != nil {
		util.PrintLog(err.Error())
		return
	}
	//然后到数据库里查询搜索结果
	// mssql := NewMssql()
	// err = mssql.Open()
	// if err != nil {
	// 	util.PrintLog(err)
	// 	return
	// }
	mssql.selectCompany()
	//sql 语句
	var historySql = "select qsrq, jzrq, tjsj, pjjs, pjzd, dds, hlcs, zms, zhgms, zmj, zhgmj, tjcs, zzl, hgzl, zxbmj, zxbzl, xbbl from schzb where bzbh='" + responseJson.Class + "' and rq='" + responseJson.Time + "'"
	mssql.queryHistoryData(historySql, responseJson.Group)
}

func getHistoryOther(otherStruct util.OtherStruct) string {
	var other string
	other = other + "开工时间=" + otherStruct.Qsrq + "&"
	other = other + "完成时间=" + otherStruct.Jzrq + "&"
	other = other + "停机时间=" + otherStruct.Tjsj + "&"
	other = other + "平均车速=" + otherStruct.Pjjs + "&"
	other = other + "平均门幅=" + otherStruct.Pjzd + "&"
	other = other + "订单笔数=" + otherStruct.Dds + "&"
	other = other + "换楞次数=" + otherStruct.Hlcs + "&"
	other = other + "累计米数=" + otherStruct.Zms + "&"
	other = other + "合格米数=" + otherStruct.Zhgms + "&"
	other = other + "累计面积=" + otherStruct.Zmj + "&"
	other = other + "合格面积=" + otherStruct.Zhgmj + "&"
	other = other + "停机次数=" + otherStruct.Tjcs + "&"
	other = other + "累计重量=" + otherStruct.Zzl + "&"
	other = other + "合格重量=" + otherStruct.Hgzl + "&"
	other = other + "累计修边(m2)=" + otherStruct.Zxbmj + "&"
	other = other + "累计修边(kg)=" + otherStruct.Zxbzl + "&"
	other = other + "修边比例=" + otherStruct.Xbbl
	util.PrintLog("History other:", other)
	return other
}
