// UdpUtils
package util

import (
	"net"
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
		PrintLog("readFromUDP, msg:", msg)
		f(msg)
	}

}
