/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/13 7:32
 */
package main

import (
	"fmt"
	"go_ex/logagent/conf"
	"go_ex/logagent/etcd"
	"go_ex/logagent/kafka"
	"go_ex/logagent/utils"
	"gopkg.in/ini.v1"
	"time"
)

// logAgent入口程序
var (
	cfg = new(conf.AppConf)
)

//func run(topic string) {
//	// 1.读取日志
//	for {
//		select {
//		case line := <-taillog.ReadChan():
//			// 2.发送到kafka
//			kafka.SendToKafka(topic, line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//
//	}
//
//}

func main() {
	// 0. 加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	// 1.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success.")

	// 2.初始化ETCD
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n", err)
		return
	}

	// 为了实现每个logagent都拉去自己独有的配置,所以要以自己的IP地址作为区分
	ipStr, err := utils.GetOutboundIP()
}
