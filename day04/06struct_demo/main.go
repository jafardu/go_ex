/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/30 20:01
 */
package main

import "fmt"

type x struct {
	a, b, c int8
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))
}
