/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/31 9:30
 */
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Go连接MySql示例

var db *sql.DB // 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	// 格式: "user:passwd@tcp(ip:port)/database_name"
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping() // 尝试连接数据库
	if err != nil {
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

// 查询单个记录
func queryOne(id int) {
	var u1 user
	// 1.查询单挑记录的sql语句
	sqlStr := `select id,name,age from user where id=?;`
	// 2.执行并拿到结果
	// 必须对rowObj对象调用Scan方法,因为该方法会释放数据库连接 // 从连接池里拿一个连接出来去数据库查询单条记录
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	// 打印结果
	fmt.Printf("u1:%#v\n", u1)
}

// 查询多条
func queryMore(n int) {
	// 1.SQL语句
	sqlStr := `select id,name,age from user where id>?;`
	// 2.执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sqlStr, err)
		return
	}
	// 3.一定要关闭rows
	defer rows.Close()

	// 4.循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed,err: %#v\n", err)
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}

// 插入数据
func insert() {
	// 1.写sql语句
	sqlStr := `insert into user(name,age) values("图朝阳",28)`
	// 2.exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 如果是插入数据的操作,能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

// 更新操作

func updateRow(newAge int, id int) {
	sqlStr := `update user set age=? where id>?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

// 删除
func deleteRow(id int) {
	sqlStr := `delete FROM user WHERE id<?;`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("删除了%d行数据\n", n)
}

// 预处理方式插入多条数据
func prepareInsert() {
	sqlStr := `INSERT INTO user(name,age) VALUES(?,?);`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("预处理失败,err:%v\n", err)
		return
	}

	defer stmt.Close()
	// 后续只需要拿到stmt去执行一些操作
	var m = map[string]int{
		"小六":  30,
		"王大锤": 35,
		"李大开": 340,
		"童绅":  38,
	}
	for k, v := range m {
		stmt.Exec(k, v) // 后续只需要传值
	}
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err: %v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
	queryOne(2)
	queryMore(1)
	//insert()
	//updateRow(9000,2)
	//for i:=3;i<=7;i++{
	//	updateRowName("xshell-0"+strconv.Itoa(i),i)
	//}
	//deleteRow(7)
	prepareInsert()
}
