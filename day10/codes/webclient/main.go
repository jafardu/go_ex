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
	"os"
)

func main() {
	response, err := http.Get("http://localhost:8888?x=1&y=2")
	fmt.Println(err)
	io.Copy(os.Stdout, response.Body)

	response, err = http.Head("http://localhost:8888?x=1&y=2")
	fmt.Println(err)
	io.Copy(os.Stdout, response.Body)
}
