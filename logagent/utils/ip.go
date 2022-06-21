/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/21 7:25
 */
package utils

import (
	"log"
	"net"
	"strings"
)

func GetOutboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "100.200.0.179:2379")
	if err != nil {
		log.Fatal("net.Dial failed!")
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr)
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}
