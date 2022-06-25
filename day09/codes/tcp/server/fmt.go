/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/25 7:49
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Printf("%05d\n", 1) //00001
	prefix := fmt.Sprintf("%05d", 1)
	ctx := []byte(prefix)

	lengthTExt := string(ctx)
	length, err := strconv.Atoi(lengthTExt)
	if err != nil {
		return
	}
	fmt.Println(length, err)
}
