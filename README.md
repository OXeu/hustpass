# HUST PASS
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