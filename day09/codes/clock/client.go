/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/26 16:01
 */
package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	protocol := "tcp"
	address := "127.0.0.1:8888"
	start := time.Now()
	conn, err := net.Dial(protocol, address)
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(conn)
	fmt.Println(reader.ReadString('\n'))
	conn.Close()
	fmt.Println("耗时:", time.Now().Sub(start))
}
