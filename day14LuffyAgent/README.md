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

GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o agentx_windows.exe

## 定时任务

1. 技术栈

etcd v3.4  (zookeeper)

robfig/cron/v3

```go
server  --- kafka
             |
             |
        - 100台agent
        - agent1
        - agent2
        ...
        - agent n
```

2. 为什么agent需要etcd或者zookeeper等类似的服务

分布式应用程序协调服务，用于管理分布式环境中的数据

方案1； server与agent直连
```go
    存在的问题：
    1. agent挂掉，server连接数增加，压力大
    2. 容错小
    3. agent没有感知能力
```

方案2：kafka作为队列
```go
    存在的问题
    1. 不能保证时效性
    2. kafka压力增加 
    3. agent挂掉，sever如何知道
```


方案3：服务中心的watch机制（服务发现）
```go
    server把任务发送给etcd
    agent与etcd建立长链接，监听etcd
    这个任务是该agent的，然后立即执行
```
    
    
健康检查
```go
- 被动
    请求接口，判断agent是否正常
- 主动
    租约id，每隔5分钟，续租
    监控服务器从etcd里面知道agent挂掉了
```
  


etcd vs zookeeper
```go
1. go开发，语言没有坑
2. 架构简单，自开发自运维
```

3. 思考

    - 如何保证任务调度的性能和实时性
    - 如何高效的监控agent以及其任务
    
4. 如何设计etcd

用途

- 用于运行和停止任务

```go
key: /cron/jobs/x.x.x.x./
value: 
	{
	"cmd": "/tmp/demo.sh",
	"express": "* * * * *",
	"task_id": task_id,
	// "type": "schedule",
	"status": "running", // pause
	"executor": "bash",
    }
```

- 存储执行任务ID

```go
key: /cron/cronid/jobs/x.x.x.x/<taskId>
value:
	{
	cronid: xxxx
    }
```

- 健康检查

```go
key: /agent/health/x.x.x.x

value: 
	{
	"status": "Online",
	"version": "v.1.1.0"
    }
```

5. etcd相关命令


- 启动

```
etcd --listen-client-urls 'http://0.0.0.0:2379' --advertise-client-urls 'http://0.0.0.0:2379'

```

- 查询kv

```
  ./etcdctl get path --from-key
```  

- 删除key

```
  ./etcdctl del path --from-key
```  

google中，etcd服务发现和租约
实际情况：etcd会出现内存泄露
原因：方法不对，短连接
      etcd内存爆了，每增加1台agent


retry-go
- 支持熔断，降级
- 重试，指数增量算法(2,4,8,...  )
