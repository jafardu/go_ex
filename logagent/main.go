/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/13 7:32
 */
package main

import (
	"fmt"
	"go_ex/logagent/conf"
	"go_ex/logagent/kafka"
	"go_ex/logagent/taillog"
	"gopkg.in/ini.v1"
	"time"
)

func run(topic string) {
	// 1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2.发送到kafka
			kafka.SendToKafka(topic, line.Text)
		default:
			time.Sleep(time.Second)
		}

	}

}

func main() {
	// 0.init conf
	cfg, err := ini.Load("conf/config.ini")
	if err != nil {
		fmt.Printf("init conf filename failed,err:%v\n", err)
		return
	}
	kconf := new(conf.AppConf)
	err = cfg.MapTo(kconf)
	if err != nil {
		fmt.Printf("macp to failed,err:%v\n", err)
		return
	}
	// 1.初始化kafka连接
	err = kafka.Init([]string{kconf.Address}, 2)
	if err != nil {
		fmt.Printf("init kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("log kafka success")
	// 2.打开日志文件准备收集日志
	err = taillog.Init(kconf.FileName)
	if err != nil {
		fmt.Printf("init taillog failed,err :%v\n", err)
		return
	}
	// 3.
	run(kconf.Topic)
}
