/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/28 15:06
 */
package main

import "fmt"

func main() {
	var a int
	a = 100
	b := &a
	fmt.Printf("type a:%T type b:%T\n", a, b)
	// 将a的十六进制内存地址打印出来
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", b) // b的值
	fmt.Printf("%v\n", b)
	fmt.Printf("%p\n", &b)
}
