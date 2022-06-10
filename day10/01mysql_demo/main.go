/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/31 7:21
 */
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
环境准备
docker run -itd --name mysql-5.7  -p 3306:3306 -e MYSQL_ROOT_PASSWORD="123456" mysql:5.7
mysql -uroot -p
create database goday10;
grant all privileges on *.* to 'root@%' identified by '123456' with grant option;
flush privileges;
*/
// Go连接Mysql示例
func main() {
	// 数据库信息
	// 用户名:密码@tcp(ip:port)/数据库名字
	dsn := "root:123456@tcp(192.168.0.xx:3306)/goday10"
	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dns:%s invalid, err:%v\n", dsn, err)
		return
	}
	err = db.Ping() // 尝试连接数据库
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功!")
}
