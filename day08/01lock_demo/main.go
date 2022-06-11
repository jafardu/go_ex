/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/24 10:03
 */
package main

import (
	"fmt"
	"sync"
)

// 锁

var (
	wg   sync.WaitGroup
	lock sync.Mutex
	x    = 0
)

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
