/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/1 14:31
 */
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// redis
var redisdb *redis.Client
var ctx = context.Background()

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:16379",
		Password: "redispass",
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err = redisdb.Ping(ctx).Result()
	return err
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed, err:%v\n", err)
		return
	}
	fmt.Println("连接redis成功!")
	// zset
	key := "rank"
	items := []*redis.Z{
		&redis.Z{Score: 90, Member: "PHP"},
		&redis.Z{Score: 96, Member: "Golang"},
		&redis.Z{Score: 97, Member: "Python"},
		&redis.Z{Score: 99, Member: "Java"},
	}
	// 把元素都追加到key
	redisdb.ZAdd(ctx, key, items...)

	// 给Golang + 10分
	newScore, err := redisdb.ZIncrBy(ctx, key, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)
}
