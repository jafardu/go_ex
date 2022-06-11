/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/28 14:53
 */
package main

import "fmt"

// 自定义类型和类型别名
type myInt int     // 自定义类型
type youtInt = int // 类型别名

func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T %d\n", n, n)

	var m youtInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
