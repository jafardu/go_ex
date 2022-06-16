/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/12 19:16
 */
package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

// tailf 的用法

func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen: true, //重新打开
		Follow: true, // 是否跟随
		Location: &tail.SeekInfo{ //从文件的那个地方开始读
			Offset: 0, Whence: 2,
		},
		MustExist: false, //文件不存在不报错
		Poll:      true,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)

	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}
}
