/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/26 18:18
 */
package main

import "net/http"

func main() {

	addr := ":8888"
	http.Handle("/w1/", http.FileServer(http.Dir("./static/")))

	http.Handle("/www/", http.StripPrefix("/www/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(addr, nil)
}
