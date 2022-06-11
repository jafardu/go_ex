/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/26 17:13
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net.http Client

// 共用一个client适用于 请求比较频繁
var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func main() {
	//resp,err := http.Get("http://127.0.0.1:20000/")
	//if err !=nil{
	//	fmt.Println("get url failed,err: ",err)
	//	return
	//}
	data := url.Values{} //url values
	urlObj, _ := url.Parse("http://127.0.0.1:20000/age/")
	data.Set("name", "jafar du")
	data.Set("age", "27")
	queryStr := data.Encode()
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("get url failed, err: ", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp.Body failed, err: %v\n", err)
		return
	}
	fmt.Println(string(b))

}
