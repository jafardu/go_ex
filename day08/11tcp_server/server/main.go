/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/27 15:39
 */
package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf) // 使用一个512字节的缓冲 `data` 来读取客户端发送来的数据并且把他们打印到服务器的终端,len 获取客户端发送的数据字节
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return // 终止程序
		}
		fmt.Printf("Received data:\n\tname:%v\n", string(buf[:len]))
	}

}
func main() {
	// 创建listen
	listen, err := net.Listen("tcp", "127.0.0.1:30001")
	if err != nil {
		fmt.Println("listen tcp failed,err: ", err)
		return
	}
	defer listen.Close()
	// listen.Accept 等待客户端的请求 客户端的请求将产生一个 net.Conn 类型的连接变量。
	for {
		conn, err := listen.Accept()
		if err != nil {
			return // 终止连接
		}
		go process(conn)

	}

}
