/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/26 16:01
 */
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	protocol := "tcp"
	address := "127.0.0.1:8888"
	listener, err := net.Listen(protocol, address)
	if err != nil {
		return
		fmt.Println(err)
	}
	fmt.Println(listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(conn)
			continue
		}
		go func() {
			fmt.Println("客户端连接成功: ", conn.RemoteAddr())
			time.Sleep(10 * time.Second)
			// 发送时间
			fmt.Fprintln(conn, time.Now().Format("2006-01-02 15:04:05"))
			conn.Close()
			fmt.Println("客户端退出: ", conn.RemoteAddr())
		}()
	}
}
