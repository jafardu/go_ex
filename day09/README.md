## 
```bash

# 测试覆盖率
go test -v
go test -v -run=TestSplit/case_01
go test -cover -coverprofile=c
go tool cover -html=c

# go tool pprof命令
go build main.go
.\main.exe -cpu=true 
go tool pprof .\cpu.pprof
# top 4 查看cpu占用情况
# list 函数名 查看函数中使用最多的代码逻辑
```

### 面试题

```shell
#台阶的走法,一次可以走1个台阶或者两个台阶
```