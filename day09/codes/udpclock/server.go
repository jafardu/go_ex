/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/26 16:34
 */
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	address := ":8888"
	protocol := "udp"

	packetConn, err := net.ListenPacket(protocol, address)
	if err != nil {
		fmt.Println(packetConn)
		return
	}
	for {
		ctx := make([]byte, 1024)
		n, addr, err := packetConn.ReadFrom(ctx)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("客户端[%s]发送数据:%s\n", addr, string(ctx[:n]))
		fmt.Println(packetConn.WriteTo([]byte(time.Now().Format("2006/01/02 15:04:05")), addr))
	}
	packetConn.Close()
}
