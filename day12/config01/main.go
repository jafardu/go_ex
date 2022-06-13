/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/10 16:22
 */
package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

type Note struct {
	Content string
	Cities  []string
}

type Person struct {
	Name string
	Age  int `ini:"age"`
	Male bool
	Born time.Time
	Note
	Created time.Time `ini:"-"`
}

type Embeded struct {
	Dates  []time.Time `delim:"|" comment:"Time data"`
	Places []string    `ini:"places,omitempty"`
	None   []int       `ini:"omitempty"`
}

type Author struct {
	Name      string `ini:"NAME"`
	Male      bool
	Age       int `comment:"Author's age"`
	GPA       float64
	NeverMind string `ini:"-"`
	*Embeded  `comment:"Embeded section"`
}

func main() {
	cfg, err := ini.Load("conf/conf.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	p := new(Person)
	err = cfg.MapTo(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Born, p.Age, p.Male)

	n := new(Note)
	err = cfg.Section("Note").MapTo(n)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("结构反射,还原配置文件")
	a := &Author{"Unknwon", true, 21, 2.8, "",
		&Embeded{
			[]time.Time{time.Now(), time.Now()},
			[]string{"Hangzhou", "Boston"},
			[]int{},
		}}
	cf := ini.Empty()
	err = ini.ReflectFrom(cf, a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.Age)
}
