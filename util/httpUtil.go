// httpUtil
package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

/*
消息格式
*/
type NormarData struct {
	Cname      string `json:"Cname"`
	Data       string `json:"Data"`
	StartTime  string `json:"StartTime"`  //开始时间
	FinishTime string `json:"FinishTime"` //完工时间
	Group      string `json:"Group"`      //产线
}

/*
订单数据
*/
type Order struct {
	Scxh        string `json:"Scxh"`   //序号
	Mxbh        string `json:"Mxbh"`   //订单号
	Khjc        string `json:"Khjc"`   //客户简称
	Zbdh        string `json:"Zbdh"`   //材质
	Klzhdh      string `json:"Klzhdh"` //楞别
	Zd          string `json:"Zd"`     //纸度
	Zbcd        string `json:"Zbcd"`   //切长
	Pscl        string `json:"Pscl"`   //排产数量
	Ddms        string `json:"Ddms"`   //留言
	Zt          string `json:"Zt"`     //是否正在进行
	Ks          string `json:"Ks"`     //剖
	Sm2         string `json:"Sm2"`
	Zbcd2       string `json:"Zbcd2"` //切长
	Xbmm        string `json:"Xbmm"`
	Scbh        string `json:"Scbh"`
	Ms          string `json:"Ms"`
	Finish_time string `json:"FinishTime"` //预计完工时间
}

/*
完工数据
*/
type FinishInfo struct {
	Mxbh     string `json:"Mxbh"`     //订单号
	Khjc     string `json:"Khjc"`     //客户简称
	Ms       string `json:"Ms"`       //米数
	Bzbh     string `json:"Bzbh"`     //班组
	Zbmc     string `json:"Zbmd"`     //材质
	Klzhdh   string `json:"Klzhdh"`   //楞别
	Zd       string `json:"Zd"`       //门幅
	Pcsl     string `json:"Pcsl"`     //排产数
	Hgpsl    string `json:"Hgpsl"`    //合格数
	Blpsl    string `json:"Blpsl"`    //不良数
	Xbmm     string `json:"Xbmm"`     //修边
	Zbcd2    string `json:"Zbcd2"`    //切长
	Ks       string `json:"Ks"`       //板宽
	Js       string `json:"Js"`       //机速
	StopTime string `json:"StopTime"` //停时
	StopSpec string `json:"StopSpec"` //停次
	Ys       string `json:"Ys"`       //用时
	Shl      string `json:"Shl"`      //损耗
}

/*
发送订单
*/
func PostOrder(data string) {
	PrintLog("post order")
	post(GetOrderUrl(), data)
}

/*
发送完工资料
*/
func PostFinihInfo(data string) {
	PrintLog("post finish_info")
	post(GetFinishInfoUrl(), data)
}

/*
获取搜索内容
*/
func GetSearchRequest(cname string) string {
	PrintLog("get search request")
	searchUrl := GetSearchRequestUrl() + "?cname=" + cname
	return get(searchUrl)
}

/*
发送搜索结果
*/
func PostSearchResult(data string) {
	PrintLog("post search result")
	post(GetSearchResultUrl(), data)
}

/*
post数据
*/
func post(httpUrl string, data string) {
	body := bytes.NewBuffer([]byte(data))
	PrintLog("body:", body)
	PrintLog("httpUrl:", httpUrl)
	resp, err := http.Post(httpUrl, HTTP_APPLICATION, body)
	if err != nil {
		PrintLog("post err:", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		PrintLog("resp err:", err)
	}
	PrintLog("resp:", string(respBody))
}

/*
get数据
*/
func get(httpUrl string) string {
	PrintLog("httpUrl:", httpUrl)
	resp, err := http.Get(httpUrl)
	if err != nil {
		PrintLog("get search request err:", err)
		return err.Error()
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		PrintLog("resp err:", err)
	}
	PrintLog("resp:", string(respBody))
	return string(respBody)
}
