/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/28 14:50
 */
package main

import "fmt"

// 递归:函数自己调用自己!
// 递归适合处理那种问题相同\问题的规模越来越小的场景
// 递归一定要有一个明确的退出条件

// 计算n的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)

}

// 上台阶的面试题
// n个台阶,一次可以走1步,也可以走两步,有多少种走法.

func f2(n uint64) uint64 {

	// 剩两个台阶 两种走法
	// 剩一个台阶 一种走法
	if n == 2 {
		return 2
	} else if n == 1 {
		return 1
	}
	return f2(n-1) + f2(n-2)
}

func main() {
	fmt.Println(f(8))
	fmt.Println(f2(8))
}
