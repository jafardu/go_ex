/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/26 19:03
 */
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 模拟浏览器
	response, err := http.Get("http://localhost:8888/help")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response.Proto)
	fmt.Println(response.StatusCode)
	io.Copy(os.Stdout, response.Body)

	req, err := http.NewRequest("DELETE", "http://localhost:8888", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	response, err = client.Do(req)

	io.Copy(os.Stdout, response.Body)

	response, err = http.Get("http://www.baidu.com")
	io.Copy(os.Stdout, response.Body)
}
