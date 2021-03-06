/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/25 7:29
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

func input(prompt string) string {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
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

	for {
		// 2.交换数据
		txt := input("请输入信息:")
		if txt == "exit" {
			break
		}
		err := Write(conn, txt)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		txt, err = read(conn)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Printf("服务器端响应:%s\n", txt)
	}
	// 3.关闭
	conn.Close()
}
