/*
* @Description:
* @Author: jafar du os4top16
* @Date: 2022/5/24 21:40

 */
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 1.与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dia 127.0.0.1:20000 failed,err:", err)
		return
	}

	// 2.发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请说话")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "text" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()

}
