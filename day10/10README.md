
### mysql
`data/sql`原生支持连接池,是并发安全的.  
这个标准库没有具体的实现,只是列出了一些需要第三方库实现的具体内容

```shell
# 环境准备
docker run -itd --name mysql-5.7  -p 3306:3306 -e MYSQL_ROOT_PASSWORD="123456" mysql:5.7
mysql -uroot -p
create database goday10;
grant all privileges on *.* to 'root@%' identified by '123456' with grant option;
flush privileges;
# 下载驱动
go get -u github.com/go-sql-driver/mysql
```

#### MySQL预处理

```text
普通SQL语句执行过程:
1.客户端对SQL语句进行占位符替换得到完整的SQL语句.
2.客户端发送完整SQL语句到MySQL服务端
3.MySQL服务端执行完整的sql语句并将结果返回给客户端.

预处理执行过程:
1.把SQL语句分成两部分,命令部分与数据部分.
2.先把命令部分发送到MySQL服务端,MySQL服务端进行SQL预处理.
3.然后把数据部分发送给MySQL服务端,MySQL服务端对sql语句进行占位符替换.
4.MySQL服务端执行完整的SQL语句并将结果返回给客户端

```

```go
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
```

### nsql

```shell
# 下载安装包 https://github.com/nsqio/nsq
nsqlookupd
nsqlookupd   -http-address=127.0.0.1:4261  -tcp-address=127.0.0.1:4260

nsqd --lookupd-tcp-address=127.0.0.1:4160   --lookupd-tcp-address=127.0.0.1:4260
nsqd -tcp-address=127.0.0.1:4250 -http-address=127.0.0.1:4251   --lookupd-tcp-address=127.0.0.1:4160  --lookupd-tcp-address=127.0.0.1:4260
nsqd -tcp-address=127.0.0.1:4350 -http-address=127.0.0.1:4351   --lookupd-tcp-address=127.0.0.1:4160 --lookupd-tcp-address=127.0.0.1:4260


nsqadmin  --lookupd-http-address=127.0.0.1:4161 --lookupd-http-address=127.0.0.1:4261

```