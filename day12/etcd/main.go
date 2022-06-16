/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/16 11:25
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
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		return
	}
	defer cli.Close()
	fmt.Println("etcd connect success")

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	_, err = cli.Put(ctx, "foo", "k8s")
	cancel()
	if err != nil {
		log.Fatal("put failed")
	}
	fmt.Println("put data success")

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "001")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}

}
