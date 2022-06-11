/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/25 12:00
 */
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
tcp 客户端
- 与服务端建立连接
- 读写数据
- 关闭
*/
func main() {

	c, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("net.Dial faild,err: ", err)
		return
	}
	// 关闭连接
	defer c.Close()

	// 键入数据
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 读取用户输入
		input, _ := inputReader.ReadString('\n')
		// 截断
		inputInfo := strings.Trim(input, "\r\n")
		// 读取到用户输入q 或者Q就退出
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		// 将输入的数据发送给服务端
		_, err := c.Write([]byte(inputInfo))
		if err != nil {
			return
		}

		buf := [512]byte{}
		n, err := c.Read(buf[:])
		if err != nil {
			fmt.Println("c.Read failed,err: ", err)
			return
		}
		fmt.Println(string(buf[:n]))

	}
}
