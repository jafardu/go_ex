/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/20 11:10
 */
package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"100.200.0.179:2379"},
		DialTimeout: time.Second,
	})
	if err != nil {
		log.Fatal("etcd connect failed!")
		return
	}

	fmt.Println("etcd connect success!")
	defer cli.Close()
	ch := cli.Watch(context.Background(), "/xxx")
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v key:%v value:%v\n ", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
		}
	}
}
