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
		log.Fatal("client v3 New failed!")
		return
	}
	fmt.Println("connect etcd success!")
	defer cli.Close()

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"c:/tmp/nginx.log","topic":"web_log"},{"path":"d:/xxx/redis.log","topic":"redis_log"}]`
	_, err = cli.Put(ctx, "/xxx", value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed,err:%v\n", err)
		return
	}

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/xxx")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%s\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("ev.Key:%s ev.value:%s\n", ev.Key, ev.Value)
	}

}
