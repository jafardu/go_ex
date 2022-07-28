/*
* @Description:
* @Author: jafar du
* @Date: 2022/7/21 7:18
 */
package main

import (
	"io"
	"net/http"
	"os"
)

/*
var html = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8" />
		<title>我的第一个页面</title>
	</head>
<body>
	<!-- 我是注释--> 我叫jafar du
	<form action="" method="GET">
		<label>姓名</label><input name="username" type="text"/>
		<label>密码</label><input name="password" type="password"/>
		<label>性别</label>
			<input type="radio" name="sex" value="1" checked="checked"/><label>男</label>
			<input type="radio" name="sex" value="0"/><label>女</label> </br>
		<label>爱好</label>
			<input type="checkbox" name="hobby" value="1" checked="checked"/><label>足球</label>
			<input type="checkbox" name="hobby" value="2"/><label>篮球</label>
			<input type="checkbox" name="hobby" value="3"/><label>羽毛球</label>
			<input type="checkbox" name="hobby" value="3"/><label>乒乓球</label><br/>
		<label>部门</label>
		<select name="department">
			<option value="dev">开发</option>
			<option value="test" selected="selected">测试</option>
			<option value="ops">运维</option>
		</select></br>
		<label>备注</label>
			<textarea name="remark"></textarea>

		<input type="submit" value="登录" /><br/>
		<input type="button" value="button" /><br/>
		<button type="button">按钮</button>
		<button type="submit">按钮</button>
	</form>
	<div>块1</div>
	<div>块2</div>
	<span>span1</span>
	<span>span2</span>
	<table>
		<thead>
			<tr>
				<th>ID</th>
				<th>姓名</th>
				<th>位置</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td>01</td>
				<td>张三</td>
				<td>上海</td>
			</tr>
			<tr>
				<td>02</td>
				<td>王五</td>
				<td>成都</td>
			</tr>
			<tr>
				<td>03</td>
				<td>os4top16</td>
				<td>甘肃</td>
			</tr>
		</tbody>
	</table>
	<ol>
		<li>洗衣服</li>
		<li>做饭</li>
		<li>看电视</li>
	</ol>
	<ul>
		<li>洗衣服</li>
		<li>做饭</li>
		<li>看电视</li>
	</ul>

	<h1>1级</h1>
	<h2>2级</h2>
	<h3>3级</h3>
	<h4>4级</h4>
	<h5>5级</h5>
	<h6>6级</h6>
	<p>第一段第一段第一段第一段第一段第一段第一段第一段第一段第一段第一段第一段第一段第一段第一段第一段</p>
	<p>第二段第二段第二段第二段第二段第二段第二段第二段第二段第二段第二段第二段第二段第二段第二段第二段</p>
	<!-- target="_blank" 在新的页面打开-->
	<a href="https://www.baidu.com" target="_blank">百度</a>
	<img src="https://img2.baidu.com/it/u=4061019804,1277113212&fm=253&fmt=auto&app=138&f=JPEG?w=347&h=500" />

</body>

</html>
`
*/
func main() {
	addr := ":8080"
	// 动态的生产响应结果
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, _ := os.Open("./template/index.html")
		io.Copy(w, f)
		//fmt.Fprint(w, html)
	})
	http.ListenAndServe(addr, nil)
}
