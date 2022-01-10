# Luffly Agent

##功能：
### 第一阶段功能：
1. Agent框架
2. 日志服务
3. 系统配置
4. 如何优雅的获取本机IP
5. http服务
6. 数据采集&设计模式

## 如何使用：

- 命令行工具

./agent start | stop | version

## 机制说明
1. 停止&启动

主要参考nginx和mysql机制；

当启动时创建pid文件；当停止时删除pid文件

2. 编译

Cross-Compilation

|OS|ARCH|OS Version|  
|------|---|---|
|linux|386 / amd63/arm| \>=Linux 2.6|
|darwin|386 / amd64| OS X (Snow Leopard + Lion)|
|freebsd|386 / amd64| \>=FreeBSD 7|
|windows|386 / amd64| \>=Windows 2000|

GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o agentx_osx

GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o agentx_linux