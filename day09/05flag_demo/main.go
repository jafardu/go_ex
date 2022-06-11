/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/27 22:34
 */
package main

import (
	"flag"
	"fmt"
)

//
func main() {
	name := flag.String("name", "jafar du", "名字")
	age := flag.Int("age", 27, "年龄")
	married := flag.Bool("married", false, "结婚了吗?")
	cTime := flag.Duration("ct", 0, "结婚多久了?")

	var name1 string
	flag.StringVar(&name1, "name1", "jafar du", "名字")

	// 先解析再使用
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)
	fmt.Println(name1)

	fmt.Println(flag.Args())  //返回命令行参数的其他参数,以[]string类型
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) //返回使用的命令行参数个数
}
