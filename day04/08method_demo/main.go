/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/30 20:14
 */
package main

import "fmt"

// 方法

// 标识符:变量名 函数名 类型名 方法名
// Go语言中如果标识符首字母是大写的,就表示对外部包可见(暴露的,公有的)

// dog这是一个狗的结构体
type dog struct {
	name string
}
type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接收者表示的是调用该方法的具体类型变量,多用类型首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪~\n", d.name)
}

// 使用值接收者:传拷贝进去

func (p person) guonian() {
	p.age++
}

// 指针接收者: 传内存地址进去
func (p *person) zhengguonian() {
	p.age++
}
func (p *person) dream() {
	fmt.Println("不上班也能挣钱!")
}
func main() {
	d1 := newDog("zhoulin")
	d1.wang()
	p1 := newPerson("元帅", 18)

	fmt.Println(p1.age)
	//p1.guonian()
	//fmt.Println(p1.age)
	p1.zhengguonian()
	fmt.Println(p1.age)
	p1.dream()
}
