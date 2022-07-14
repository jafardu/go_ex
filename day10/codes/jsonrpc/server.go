/*
* @Description:
* @Author: jafar du
* @Date: 2022/7/11 7:57
 */
package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type AddRequest struct {
	Left, Rright int
}

type AddRespone struct {
	Result int
}

type Calc struct{}

// 参数1: 请求对象(可以是指针/值)
// 参数2: 响应对象(可以是指针)
// 返回值: error

func (c *Calc) Add(req AddRequest, resp AddRespone) error {
	fmt.Println("calc.add")
	time.Sleep(10 * time.Second)
	resp.Result = req.Rright + req.Left
	return nil
}

func main() {
	// 注册RPC服务
	rpc.Register(&Calc{})
	rpc.RegisterName("AAAA", &Calc{})

	listener, _ := net.Listen("tcp", "0.0.0.0:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		go jsonrpc.ServeConn(conn)
	}
	listener.Close()
}
