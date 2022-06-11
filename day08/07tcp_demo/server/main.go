/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/24 21:47
 */
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp server端
func processConn(conn net.Conn) {
	defer conn.Close()
	// 3.与客户端通信
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed,err", err)
			return
		}
		fmt.Println(string(tmp[:n]))
		fmt.Print("请回复: ")
		msg, _ := reader.ReadString('\n') // 读到换行
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
}
func main() {
	// 1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("open listen failed,err:", err)
		return
	}
	// 2.等待别人跟自己建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed,err", err)
			return
		}
		// 使用goroutine去处理连接
		go processConn(conn)
	}

}
