/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/20 11:49
 */
package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// GetOutboundIP 获取本地对外IP地址
func GetOutbounIP() (ip string, err error) {
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
func main() {
	ip, err := GetOutbounIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}
