/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/13 7:32
 */
package main

import (
	"fmt"
	"go_ex/logagent/kafka"
	"go_ex/logagent/taillog"
	"time"
)

func run() {
	// 1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2.发送到kafka
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}

	}

}

func main() {
	// 1.初始化kafka连接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("init kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("log kafka success")
	// 2.打开日志文件准备收集日志
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("init taillog failed,err :%v\n", err)
		return
	}
	// 3.
	run()
}
