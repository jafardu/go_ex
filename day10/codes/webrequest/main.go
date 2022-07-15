/*
* @Description:
* @Author: jafar du
* @Date: 2022/7/15 7:23
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := "0.0.0.0:8888"
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// 请求行
		fmt.Println("method:", r.Method)
		fmt.Println("URL:", r.URL)
		fmt.Println("protocol:", r.Proto)

		// 请求头
		h := r.Header
		fmt.Println(h.Get("User-Agent"))

		fmt.Println(h.Get("Token"))

		fmt.Fprint(w, time.Now().Format("2006-01+02 15:04:05"))

		// 请求体
		fmt.Println("body:")
		io.Copy(os.Stdout, r.Body)

	})
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err)
}
