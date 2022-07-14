/*
* @Description:
* @Author: jafar du
* @Date: 2022/7/12 7:31
 */
package main

import (
	"net"
	"net/rpc"
)

type AddRequest struct {
	left  int
	right int
}

type AddResponse struct {
	result int
}

type Calc struct{}

func (c *Calc) Add(req AddRequest, resp AddResponse) error {
	resp.result = req.right + req.left
	return nil
}

func main() {

	//rpc
	// 注册
	rpc.Register(&Calc{})
	//启动服务
	listenenr, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		return
	}
	rpc.Accept(listenenr)
	listenenr.Close()

}
