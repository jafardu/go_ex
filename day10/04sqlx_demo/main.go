/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/1 11:35
 */
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
	"github.com/jmoiron/sqlx"
)

// Go连接MySQL示例

var db *sqlx.DB // 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	dsn := `root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True`
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	ID   int
	NAME string
	AGE  int
}

// CRUD
// 查询单条数据示例
func queryRowDemo() {
	sqlStr := `select id,name,age from user where id=?`
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.NAME, u.AGE)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := `select id,name,age from user where id>?`
	var usrs []user
	err := db.Select(&usrs, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", usrs)
}

// 更新数据
func updateRowDemo() {
	sqlStr := `update user set age=? where id=?`
	ret, err := db.Exec(sqlStr, 9000, 1)
	if err != nil {
		fmt.Printf("update failed, err:%#v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%#v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("delete success,affected rows:%d\n", n)
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%#v\n", err)
		return
	}
	sqlStr1 := `select id,name,age from user where id=9`
	var u user
	db.Get(&u, sqlStr1)

	fmt.Printf("u:%#v\n", u)

	var userList []user
	sqlStr2 := `select id,name,age from user`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return
	}
	fmt.Printf("userList:%#v\n", userList)

	fmt.Println("查询单条数据")
	queryRowDemo()
	fmt.Println("查询多行数据")
	queryMultiRowDemo()
	fmt.Println("插入数据")
	fmt.Println("更新数据")
	updateRowDemo()
	fmt.Println("删除数据")
}
