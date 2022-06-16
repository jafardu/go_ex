/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/27 16:31
 */
package main

import (
	"fmt"
	"go_ex/day08/09nianbao_jiejue/protocol"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		log.Fatalf("failed to Dial : %v", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		//inputReader := bufio.NewReader(os.Stdin)
		//fmt.Println("请输入自己的名字:")
		//clientName,_:=inputReader.ReadString('\n')
		//trimInput:=strings.Trim(clientName,"\r\n")
		//if trimInput == "Q" || trimInput == "q" {
		//	return
		//}
		//conn.Write([]byte(trimInput+"says:" + trimInput))
		msg := `Hello, Hello. How are you?`
		// 调用协议编码数据
		b, err := protocol.Encode(msg)
		if err != nil {
			fmt.Println("encode failed,err:", err)
			return
		}
		conn.Write(b)
	}
}
