/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/25 11:37
 */
package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
tcp server流程:
监听端口
接收客户端请求建立链接
创建goroutine处理链接
关闭
*/

func process(c net.Conn) {
	// 关闭连接
	defer c.Close()
	for {
		// 创建一个用来从文件读取内容的对象
		reader := bufio.NewReader(c)
		var buf [256]byte
		// 读取数据
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("reader.Read failed,err: ", err)
			break
		}
		recvData := string(buf[:n])
		fmt.Println("receive data : ", recvData)
		// 将数据再发给客户端
		c.Write([]byte(recvData))
	}
}
func main() {
	// 监听tcp
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("net.Listen error: ", err)
		return
	}
	for {
		// 建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept failed,err: ", err)
			continue
		}
		// 专门开一个goroutine去处理连接
		go process(conn)
	}
}
