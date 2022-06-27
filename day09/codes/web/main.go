/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/26 17:53
 */
package main

import (
	"fmt"
	"net/http"
)

// Home 处理器函数
func Home(w http.ResponseWriter, r *http.Request) {
	// 用户提交的数据 http内容->go代码转换 http.Request
	w.Write([]byte("hi"))
	fmt.Println("xxxx")
}

type Help struct{}

func (h *Help) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "help")
}
func main() {
	addr := ":8888"
	// 处理器函数

	// url => 找处理器函数 => 调用处理器函数(http包)
	// 指定url和处理器关系
	// 处理器函数签名由http包定义
	http.HandleFunc("/home", Home)
	http.Handle("/help", new(Help))

	http.ListenAndServe(addr, nil)
}
