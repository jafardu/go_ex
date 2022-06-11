/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/24 10:37
 */
package main

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	wg sync.WaitGroup
	// Go内置的map不是并发安全的,超过21个并发的写入肯定报错
	m    = make(map[string]int)
	lock sync.Mutex
)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			lock.Lock()
			set(key, n)
			lock.Unlock()
			fmt.Printf("k:=%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
