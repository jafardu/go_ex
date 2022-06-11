/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/27 16:22
 */
package main

import (
	"bufio"
	"fmt"
	"go_ex/day08/09nianbao_jiejue/protocol"
	"io"
	"log"
	"net"
)
func process(conn net.Conn){
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		//buf := make([]byte,512)
		//len,err :=conn.Read(buf)
		//if err !=nil{
		//	log.Fatalf("read failed : %v", err)
		//}
		//fmt.Printf("解析数据: %v",string(buf[:len]))
		recvStr,err := protocol.Decode(reader)
		if err == io.EOF{
			return
		}
		if err !=nil{
			fmt.Println("decode failed,err:",err)
		}
		fmt.Println("收到client发来的数据:",recvStr)
	}
}
func main(){
	listen,err := net.Listen("tcp","127.0.0.1:30000")
	if err!=nil{
		fmt.Println("监听打开失败,err: ",err)
		return
	}
	//等待客户端的连接
	for {
		conn,err := listen.Accept()
		if err !=nil{
			log.Fatal(err)
		}
		go process(conn)
	}
}
