package main

import (
	"day14LuffyAgent/cron"
	"day14LuffyAgent/http"
	"day14LuffyAgent/metrics"
	"day14LuffyAgent/settings"
	"fmt"
	goHttp "net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main()  {

	ch := make(chan os.Signal, 1)

	// 捕获os信号
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		signalType := <- ch
		signal.Stop(ch)
		fmt.Println("退出")
		fmt.Println("收到的OS信号类型是", signalType)

		//  删除pid
		os.Remove(settings.Config().Pid)
		os.Exit(0)
	}()

	if len(os.Args) != 2{
		fmt.Printf("使用说明:%s [start|stop|version\n]", os.Args[0])
		os.Exit(0)   // 安全退出
	}

	settings.LoadConfiguration() // 加载配置文件
	settings.InitLocalIp()  // 知晓ip地址


	/*
		监控数据定时采集
		先初始化，引入映射关系，然后进行数据采集

		解耦：数据定义，定时，数据采集
		插拔式，容易实现服务降级，熔断。

		监控采集服务挂掉，不影响其他服务运行
	*/
	metrics.BuildMappers()



	go cron.InitDataHistory()
	cron.Collect()

	// 让agent有守护进程的能力
	fmt.Println(os.Args[1])
	fmt.Println("开始启动守护进程...")
	if strings.ToLower(os.Args[1]) == "main" {
		fmt.Println("启动中...")
		go func() {
			goHttp.ListenAndServe("0.0.0.0:16060", nil)
		}()
		http.Start()
		fmt.Println("启动完成")
	}
	settings.HandleControl(os.Args[1])

}


/*


kill -9 PID
不能被任何程序捕获，一般使用 kill -15 xxx

github.com/valyala/fasthttp
fasthttp并发80w


编译：
GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o agentx_windows.exe

运行
$ ./agentx_windows.exe start

查看进程
$ tasklist.exe |findstr 16884

访问地址
http://192.168.1.20:8080/v1/check/health


删除进程
$ taskkill.exe -PID 16884 -F
*/