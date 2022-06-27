/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/27 7:27
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func write(conn net.Conn, txt string) error {
	length := len(txt)
	_, err := conn.Write([]byte(fmt.Sprintf("%05d", length)))
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = conn.Write([]byte(txt))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
func read(conn net.Conn) (string, error) {
	lengthBytes := make([]byte, 5)
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
	return string(ctx), err
}
func input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
func main() {
	addr := "127.0.0.1:8888" //host:port
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
		// 3. 与客户端交换数据
		for {
			txt, err := read(conn)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			fmt.Printf("客户端发送: %s\n", txt)
			txt = input("请输入信息: ")
			if txt == "exit" {
				break
			}
			err = write(conn, txt)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
		}
		// 4.关闭客户端连接
		conn.Close()
	}

	// 5.关闭服务器
	listener.Close()
}
