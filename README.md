# xhttp
golang 文件服务器

## 使用
可直接下载release版本（只适用于mac），其它版本可自行构建

> xhttp -h  # 获取帮助信息


```
Usage of xhttp:
  -d string
    	监听目录绝对路径,xhttp -d /www/dist (default "/home")
  -p string
    	监听端口,xhttp -p 9001  (default "9000")
```

> xhttp -d /www # 直接在9000端口上打开/www目录浏览



> xhttp -d /www -p 80 # 在80端口上打开/www目录浏览
