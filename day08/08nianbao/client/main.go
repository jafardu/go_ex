/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/27 15:04
 */
package main

import (
	"fmt"
	"net"
)

func main(){
	conn,err := net.Dial("tcp","127.0.0.1:30000")
	if err !=nil{
		fmt.Println("连接失败,err:",err)
		return
	}
	defer conn.Close()
	for i:=0;i<20;i++{
		msg := `hello,hello. How are you?`
		conn.Write([]byte(msg))
	}

}
