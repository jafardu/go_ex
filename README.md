# go_ex

## 测试
## 工具
### goland
#### 1.快捷键
ctrl+j 快速输入模块

#### 2.常用插件
```text
gorm sql 编写补全 

Save Actions. ctrl+s 可以自动格式化代码，导包等
Gopher go仓鼠进度条
Key Promoter X 在用鼠标使用idea功能时会提示你该功能的快捷键
Rainbow Brackets 彩色的代码括号
String Manipulation 字符串操作

```
#### 3.配置
**配置file templates** 
```text
Ctrl+Alt+S添加以下内容到Go File文件
/*
* @Description:
* @Author: jafar du
* @Date: ${DATE} ${TIME}
 */
package ${PROJECT_NAME}
```
**配置git代理**
```shell
 git config --global https.proxy http://127.0.0.1:7890
 git config --global http.proxy http://127.0.0.1:7890
 
 # 取消配置
 git config --global --unset http.proxy 
 git config --global --unset https.proxy 
```

#### 4.杂记

https://www.youtube.com/watch?v=cWLDQ4C4BSc