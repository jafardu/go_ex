/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/25 7:29
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
	addr := "127.0.0.1:8888"
	protocol := "tcp"

	// 1.创建连接
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2.交换数据
	Write(conn, "nihao")
	//conn.Write([]byte("你好"))
	//conn.Read()
	// 3.关闭
	conn.Close()
}
