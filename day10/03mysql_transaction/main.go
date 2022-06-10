/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/1 10:50
 */
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	// 数据库信息
	// 用户名:密码@tcp(ip:port)/数据库的名字
	dsn := `root:123456@tcp(100.200.0.179:3306)/sql_test`
	// 连接数据库
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping() // 尝试连接数据库
	if err != nil {
		fmt.Printf("数据库连接失败,err:%s\n", err)
		return
	}
	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  // 设置最大空闲连接数
	return
}

type user struct {
	id   int
	name string
	age  int
}

func transactionDemo() {
	// 1.开启事务
	tx, err := db.Begin()
	if err != nil {
		return
	}
	// 2.执行多个SQL操作
	sqlStr1 := `update user set age=age-2 where id=7`
	sqlStr2 := `update user set age=age+2 where id=8`
	// 3.执行SQL1
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行SQL1出错,要回滚")
		return
	}
	// 4.执行SQL2
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行SQL2出错,要回滚")
		return
	}
	// 上面两步SQL都执行成功,就提交本次事务
	err = tx.Commit()
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("提交出错啦,要回滚!")
		return
	}
	fmt.Println("事务执行成功!")

}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%#v\n", err)
	}
	fmt.Println("连接数据库成功!")
	transactionDemo()
}
