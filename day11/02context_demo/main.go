/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/4 18:28
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要context?
var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

func f() {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("jafar")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break FORLOOP
		default:
		}
	}
}
func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	// 如何通知子goroutine退出
	exitChan <- true
	wg.Wait()
}
