/*
* @Description:
* @Author: jafar du
* @Date: 2022/7/11 7:51
 */
package main

import (
	"fmt"
	"net/rpc"
)

// 远程服务Add(1,2) int
type AddRequest struct {
	Left  int
	Right int
}

// type AddRespone int
type AddResponse struct {
	Result int
}

func main() {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:8888")

	req := AddRequest{3, 10}
	resp := AddResponse{}

	err := client.Call("Calc.Add", req, &resp)
	fmt.Println(err, resp)

	client.Close()
}
