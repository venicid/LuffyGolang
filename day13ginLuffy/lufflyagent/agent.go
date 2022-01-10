// author： Boyle Gu
package main

import (
	"fmt"
	"lufflyagent/cron"
	"lufflyagent/http"
	"lufflyagent/logger"
	"lufflyagent/metrics"
	"lufflyagent/settings"
	goHttp "net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main()  {
	logger.StartupDebug("begin")
	ch := make(chan os.Signal,1)

	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		signalType := <-ch
		signal.Stop(ch)
		fmt.Println("退出...")
		fmt.Println("收到OS信号类型 : ", signalType)
		logger.StartupDebug("收到OS信号类型 : ", signalType)
		// 删除pid
		os.Remove(settings.Config().Pid)
		os.Exit(0)
	}()


	if len(os.Args) != 2 {
		fmt.Printf("使用说明 : %s [start|stop|version] \n ", os.Args[0])
		os.Exit(0) // 安全退出
	}

	settings.LoadConfiguration()
	settings.InitLocalIp()
	metrics.BuildMappers()
	go cron.InitDataHistory()
	cron.Collect()

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
