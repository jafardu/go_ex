/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/29 22:38
 */
package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
