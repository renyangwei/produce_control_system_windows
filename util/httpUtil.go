// httpUtil
package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

/*
消息格式
*/
type NormarData struct {
	Cname      string    `json:"cname"`
	Data       string    `json:"data"`
	StartTime  time.Time `json:"StartTime"`  //开始时间
	FinishTime time.Time `json:"FinishTime"` //完工时间
	Group      string    `json:"Group"`      //产线
}

/*
订单数据
*/
type Order struct {
	Mxbh        string    `json:"Mxbh"`       //订单号
	Khjc        string    `json:"Khjc"`       //客户简称
	Zbdh        string    `json:"Zbdh"`       //材质
	Klzhdh      string    `json:"Klzhdh"`     //楞别
	Xdzd        string    `json:"Xdzd"`       //纸度
	Pscl        string    `json:"Pscl"`       //排产数量
	Zbcd        string    `json:Zbcd""`       //切长
	Ks          string    `json:"Ks"`         //剖
	Finish_time time.Time `json:"FinishTime"` //预计完工时间
	Ddms        string    `json:"Ddms"`       //留言
}

/*
完工数据
*/
type FinishInfo struct {
	Mxbh  string `json:"Mxbh"`  //订单号
	Khjc  string `json:"Khjc"`  //客户简称
	Zbdh  string `json:"Zbdh"`  //材质
	Zbkd  string `json:"Zbkd"`  //纸板宽
	Hgpsl string `json:"Hgpsl"` //合格数
	Blpsl string `json:"Blpsl"` //不良数
	Pcsl  string `json:"Pcsl"`  //排产数
	Zbcd  string `json:"Zbcd"`  //切长
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

func post(httpUrl string, data string) {
	body := bytes.NewBuffer([]byte(data))
	PrintLog("body:", body)
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
