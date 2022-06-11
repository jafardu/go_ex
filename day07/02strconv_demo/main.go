package main

import (
	"fmt"
	"strconv"
)

// strconv 包实现了基本数据类型和其字符串表示的相互转换
// 主要的函数包含 Atoi()、Itoa()、parse系列、format系列、append系列
func main() {
	// 从字符串中解析出对应的数据
	str := "10000"
	//ret1 := int64(str)
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parseint failed, err:", err)
		return
	}
	fmt.Printf("%#v\n", ret1, int(ret1))

	// Atoi: 字符串转换成int
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)

	// 从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)

	// 从字符串中解析出浮点数
	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue)

	// 把数字转换成字符串类型
	i := 97
	//ret2 := string(i) // "a"

	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v %T\n", ret2, ret2)
	ret3 := strconv.Itoa(i)
	fmt.Printf("%#v %T\n", ret3, ret3)

}
