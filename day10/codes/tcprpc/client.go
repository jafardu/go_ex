/*
* @Description:
* @Author: jafar du
* @Date: 2022/7/12 7:31
 */
package main

import (
	"fmt"
	"net/rpc"
)

// 远程服务 Add(1,2) int
type AddRequest struct {
	Left, Right int
}

// type AddResponse int
type AddResponse struct {
	Result int
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 调用
	req := AddRequest{1, 2}
	resp := AddResponse{}

	// func,req,resp
	err = client.Call("Calc.Add", req, &resp)
	fmt.Println(err, resp)
	client.Close()
}
