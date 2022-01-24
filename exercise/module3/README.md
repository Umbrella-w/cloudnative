README

dockerfile:

```
FROM golang:1.17 as builder
WORKDIR /go/src
ENV GOPROXY=https://goproxy.cn
COPY . .
```

构建镜像: `docker build -t umbrellacn/web1 .`



docker-compose

```yml
version: '3'
services:
  app:
    image: umbrellacn/web1:latest
    ports:
      - "80:80"
    command: go run /go/src/main.go

```

本机有安装docker-compose可以使用：

启动命令：`docker-compose up -d`

![image-20220116202908047](C:\Users\Umbre\AppData\Roaming\Typora\typora-user-images\image-20220116202908047.png)

若没有docker-compose: 

`mkdir -p /root/src && cd /root/src`

`git clone https://github.com/Umbrella-w/cloudnative.git`

`docker run -itd --name web -p 80:80 -v /root/src/cloudnative/exercise/module3/testgo:/go/src umbrellacn/web1:latest`


ip访问：http://1.13.168.134/
