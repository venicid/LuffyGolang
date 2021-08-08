package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var quitC = make(chan struct{})

func signalWork()  {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 当c中读取到值的时候，说明有人发了信号
	sig := <- c

	// 通知所有读取quitC的任务
	close(quitC)
	time.Sleep(2*time.Second)
	log.Printf("接收到了停止的信号,信号是%v，pod:=%d, 要退出了",sig, os.Getpid())
}

func work01()  {
	ticker := time.NewTicker(5*time.Second)
	for{
		select{
		case <- ticker.C:
			log.Printf("[我是woker01][5秒周期到了，干活]")

		case <- quitC:
			log.Printf("[我是woker01][接受到主进程退出的信号]。进行清理操作")
			return
		}
	}
}

func work02()  {
	ticker := time.NewTicker(5*time.Second)
	for{
		select{
		case <- ticker.C:
			log.Printf("[我是woker02][5秒周期到了，干活]")

		case <- quitC:
			log.Printf("[我是woker02][接受到主进程退出的信号]。进行清理操作")
			return
		}
	}
}

func work03()  {
	ticker := time.NewTicker(5*time.Second)
	for{
		select{
		case <- ticker.C:
			log.Printf("[我是woker02][5秒周期到了，干活]")

		case <- quitC:
			log.Printf("[我是woker02][接受到主进程退出的信号]。进行清理操作")
			return
		}
	}
}

func main() {

	go work01()
	go work02()
	go work03()

	signalWork()
}

/*
2021/08/08 09:45:47 [我是woker02][5秒周期到了，干活]
2021/08/08 09:45:47 [我是woker01][5秒周期到了，干活]
2021/08/08 09:45:47 [我是woker02][5秒周期到了，干活]
2021/08/08 09:45:52 [我是woker02][5秒周期到了，干活]
2021/08/08 09:45:52 [我是woker02][5秒周期到了，干活]
2021/08/08 09:45:52 [我是woker01][5秒周期到了，干活]
2021/08/08 09:45:53 [我是woker01][接受到主进程退出的信号]。进行清理操作
2021/08/08 09:45:53 [我是woker02][接受到主进程退出的信号]。进行清理操作
2021/08/08 09:45:53 [我是woker02][接受到主进程退出的信号]。进行清理操作
2021/08/08 09:45:53 接收到了停止的信号,信号是interrupt，pod:=3004, 要退出了

Process finished with the exit code 0

*/