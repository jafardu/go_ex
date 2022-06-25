/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/25 7:10
 */
package main

import (
	"fmt"
	"net"
	"strconv"
)

func Write(conn net.Conn, txt string) error {
	length := len(txt)
	n, err := conn.Write([]byte(fmt.Sprintf("%05d", length)))
	if err != nil {
		fmt.Println(n, err)
		return err
	}
	n, err = conn.Write([]byte(txt))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func read(conn net.Conn) (string, error) {
	lengthBytes := make([]byte, 5)
	//conn.Read(lengthBytes)
	_, err := conn.Read(lengthBytes)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	length, err := strconv.Atoi(string(lengthBytes))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	ctx := make([]byte, length)
	_, err = conn.Read(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(ctx), nil
}

func main() {
	addr := "127.0.0.1:8888" // host:port
	protocol := "tcp"
	// 1.监听服务
	listener, err := net.Listen(protocol, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(listener.Addr())
	// 2.接受客户端请求
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(conn.LocalAddr(), conn.RemoteAddr())
		// 3.与客户端交换数据
		fmt.Println(read(conn))
		//ctx := make([]byte, 1024)
		//n, err := conn.Read(ctx)
		//fmt.Println(string(ctx[:n]), n, err)
		//conn.Write()
		conn.Close()
	}
	// 4.关闭客户端
	// 5.关闭服务器
	listener.Close()
}
