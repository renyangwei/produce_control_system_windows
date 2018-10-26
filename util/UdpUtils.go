// UdpUtils
package util

import (
	"net"
	"strings"
)

/*
开始监听UDP
*/
func StartListenUdp(f func(a string)) {
	//获得端口
	udpPort := ParamUdpPort()
	PrintLog("udp port is " + udpPort)
	PrintLog("start listening udp...")
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+udpPort)
	if err != nil {
		PrintLog("resolveUDPAddress error:" + err.Error())
		return
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		PrintLog("listenUDP error:" + err.Error())
		return
	}

	for {
		var buf [1024]byte
		n, _, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			PrintLog("readFromUDP error:" + err.Error())
			return
		}
		msg := string(buf[0:n])
		PrintLog("readFromUDP, msg:", replaceNumber(msg))
		f(msg)
	}
}

/*
替换数字，如果有数字1，则替换为中文的一
*/
func replaceNumber(msg string) string {
	if strings.Contains(msg, "1号线") {
		return strings.Replace(msg, "1号线", "一号线", -1)
	} else if strings.Contains(msg, "2号线") {
		return strings.Replace(msg, "2号线", "二号线", -1)
	} else if strings.Contains(msg, "3号线") {
		return strings.Replace(msg, "3号线", "三号线", -1)
	} else if strings.Contains(msg, "4号线") {
		return strings.Replace(msg, "4号线", "四号线", -1)
	} else if strings.Contains(msg, "5号线") {
		return strings.Replace(msg, "5号线", "五号线", -1)
	} else if strings.Contains(msg, "6号线") {
		return strings.Replace(msg, "6号线", "六号线", -1)
	} else {
		return msg
	}

}
