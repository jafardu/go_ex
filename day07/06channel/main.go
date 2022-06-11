/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/23 18:04
* @LastEditors: jafar du
 */
package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int // 需要制定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)
	b = make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()

	b <- 10 // hang住了
	fmt.Println("10发送到通道b中了...")
	wg.Wait()
}

func bufChannel() {
	fmt.Println(b)         // nil
	b = make(chan int, 10) // 不带缓冲通道的初始化
	b <- 10                // hang 住了
	fmt.Println("10发送到通道b中了...")
	b <- 20
	fmt.Println("20发送到通道中了...")
	x := <-b
	fmt.Println("从通道b中取到了", x)
	close(b)
}

func main() {
	//noBufChannel()
	bufChannel()
}
