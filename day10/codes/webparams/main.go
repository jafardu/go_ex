/*
* @Description:
* @Author: jafar du
* @Date: 2022/7/15 8:07
 */
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	addr := "0.0.0.0:8888"
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Method)
		// 1. 提交数据方式
		/*
			在URL中传递数据
			url?argname1=argvalue1&argname2=argvalue2
		*/
		//fmt.Println(r.URL)
		r.ParseForm() // 解析参数
		fmt.Println(r.Form)
		fmt.Println(r.Form.Get("a"))
		fmt.Println(r.Form["a"])

		// 2. 通过body提交数据 request.Body Post
		/* body中数据格式
		3种
		application/x-www-form-urlencoded a=b&c=d
		multipart/form-data

		其它类型 = > 解码
		application/json
		application/xml
		自己定义
		a=b&
		*/
		/*
			a. application/x-www-form-urlencoded a=b&c=d
				Form 可以获取URL中参数也可以获取Body中参数
				curl -XPOST  "127.0.0.1:8888/port?a=1&b=2" -d"x=10&y=20"
					map[a:[1] b:[2] x:[10] y:[20]]
		*/
		fmt.Println(r.PostForm)

		fmt.Fprint(w, time.Now().Format("2006-05-04 15:04:05"))
	})
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err)
}
