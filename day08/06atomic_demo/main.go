/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/24 10:52
 */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作
var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex
)

func add() {
	//x ++
	//lock.Lock()
	//x =x+1
	//lock.Unlock()

	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)

	// 比较并交换
	x = 200
	ok := atomic.CompareAndSwapInt64(&x, 200, 300)
	fmt.Println(ok, x)
}
