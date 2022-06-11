/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/27 22:31
 */
package main

import (
	"fmt"
	"os"
)

// os.Args 获取命令行参数
func main() {

	fmt.Println(os.Args)
	fmt.Println(os.Args[0], os.Args[2])

}
