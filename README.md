# HUST PASS
[![Build Status](https://drone.kafi.work/api/badges/xeu/hust_pass/status.svg)](https://drone.kafi.work/xeu/hust_pass)

自动过校园网认证
## 使用说明
### 1.配置修改
clone 该仓库，修改 `main.go` 中的 `username` 和 `password` 常量
为自己的校园网账号密码
### 2.Windows 运行
安装 `Golang` 编译环境, 执行以下命令

```go build -ldflags "-H=windowsgui"```

`-ldflags "-H=windowsgui"` 
指定进程运行在后台,无控制台窗口,可在任务管理器中结束进程

### 3.Linux Shell运行
```nohup go run main.go &```

```nohup *** &``` 表示在后台运行,输出保存在当前目录下的
`nohup.out`文件中

### 4.Linux Docker运行
安装`Docker`以及`Docker-Compose`

在当前文件夹下执行
```docker-compose up -d```

将自动下载 构建镜像 进行构建并打包镜像自动运行

无需安装`Golang`编译环境


# 常见问题
## 1. 启用 Docker 后认证网页无法访问
大概率是Docker自动生成的网络把`172.18.18.0`这个子网段给占了导致认证请求回到了容器内  
输入 `ifconfig` 查看网络配置  
如果存在如下配置的网络
```
br-0aaf5b13ef7d ....
inet addr:172.18.0.1  Bcast:172.18.255.255(这个不重要)  Mask:255.255.0.0
```
则输入`docker  network ls` 查看 `0aaf5b13ef7d`(`br-`后的跟随的内容) 对应的虚拟网卡，使用`docker network rm 0aaf5b13ef7d` 移除该虚拟网卡，如果移除失败，需先移除使用该网卡的容器`docker ps`查看活跃容器,`docker stop (容器id)`停止容器，之后再移除网卡

