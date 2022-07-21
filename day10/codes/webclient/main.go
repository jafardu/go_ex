/*
* @Description:
* @Author: jafar du
* @Date: 2022/7/18 7:51
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	response, err := http.Get("http://localhost:8888?x=1&y=2")
	fmt.Println(err)
	io.Copy(os.Stdout, response.Body)

	response, err = http.Head("http://localhost:8888?x=1&y=2")
	fmt.Println(err)
	io.Copy(os.Stdout, response.Body)

	// x-www-form-urlencoded
	values := url.Values{}
	values.Add("x", "1")
	values.Add("x", "2")
	values.Set("y", "1")
	values.Set("y", "2")
	response, err = http.PostForm("http://locahost:8888", values)
	fmt.Println(err)
	io.Copy(os.Stdout, response.Body)

	// application/json
	f, _ := os.Open("test.json")
	response, err = http.Post("http://localhost:8888", "application/json", f)
	fmt.Println(err)
	io.Copy(os.Stdout, response.Body)

}
