/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/27 14:44
 */
package main

import (
	"fmt"
	"go_ex/day09/split_string"
)

func main() {
	ret := split_string.Split("hello,world","l")
	fmt.Printf("%#v\n",ret)
}