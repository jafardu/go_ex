## 作业

1. socket  
   远程的文件操作  
   远程服务器  
   客户端  

        client -cmd
            ls -path / 获取服务路径下的文件信息
            put -src -dst 上传文件
            get -src -dst 下载文件
            rm -path 移除文件

        请求: length{"cmd":"", path:"", content:""}
        响应: length{"status":"200", "content:""}

        content太大的时候如何办
        5: 1024 * 10 10KB
        content 10KB

        /home/kk/fileserver

        ../../..

2. 用户管理(挑战)  
   socket版本  
   用户添加  
   用户查询  

3. 压缩算法(挑战)  
   zlib  
   zip  


c/s模式开发