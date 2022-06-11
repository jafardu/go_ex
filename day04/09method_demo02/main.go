/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/30 20:28
 */
package main

import "fmt"

// 给自定义类型加方法

type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

func main() {
	m := myInt(100)
	m.hello()
}
