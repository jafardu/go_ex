/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/27 15:47
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
	// 打开连接:
	conn, err := net.Dial("tcp", "127.0.0.1:30001")
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字: ")
	clientName, _ := inputReader.ReadString('\n')
	trimedClient := strings.Trim(clientName, "\r\n")

	// 给服务器发送信息直到退出
	for {
		fmt.Println("发送消息到服务端,请输入[输入Q退出连接]:")
		input, _ := inputReader.ReadString('\n')
		trimedInput := strings.Trim(input, "\r\n")
		if trimedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimedClient + " \n\tsays: " + trimedInput))

	}

}
