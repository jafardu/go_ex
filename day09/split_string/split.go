/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/26 7:51
 */
package split_string

import (
	"strings"
)

func Split(s, sep string) []string {
	// s="abcdcfce" sep="c" [ab d f e]
	//var ret []string
	ret := make([]string, 0, strings.Count(s, sep))
	index := strings.Index(s, sep) // 2
	for index >= 0 {
		ret = append(ret, s[:index]) // [ab]
		s = s[index+len(sep):]
		index = strings.Index(s, sep) // index=1
	}
	ret = append(ret, s)
	return ret
}

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
