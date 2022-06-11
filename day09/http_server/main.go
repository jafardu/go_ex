/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/25 21:12
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./index.html")
	if err != nil {
		fmt.Println("ioutil.ReadFile failed,err: ", err)
		return
	}
	w.Write([]byte(b))
}
func f2(w http.ResponseWriter, r *http.Request) {
	// 对于Get请求,参数都放在url上(query param),请求体中是没有数据的.
	queryParam := r.URL.Query() // 自动帮我们识别URL中的query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}
func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/name/", f2)
	http.HandleFunc("/age/", f2)
	http.ListenAndServe("127.0.0.1:20000", nil)
}
