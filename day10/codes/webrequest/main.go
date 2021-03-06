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
	"strings"
	"time"
)

func main() {
	addr := "0.0.0.0:8888"
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(strings.Repeat("-", 30))

		// 请求行
		fmt.Println("method:", r.Method)
		fmt.Println("URL:", r.URL)
		fmt.Println("protocol:", r.Proto)

		// 请求头
		fmt.Println(r.Host)

		h := r.Header
		fmt.Println(h.Get("User-Agent"))

		fmt.Println(h.Get("Token"))

		// 请求体
		fmt.Println("body:")
		io.Copy(os.Stdout, r.Body)

		fmt.Fprint(w, time.Now().Format("2006-01+02 15:04:05"))

	})
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err)
}
