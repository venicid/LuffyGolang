

package main

import (
"log"
"time"
)

func main() {

	// 初始化一个类型为int的chan
	data := make(chan int, 3)

	// 达成和最后的time.sleep一样的效果 阻主线程，防止异步任务中有未处理完的就推出
	quit := make(chan bool)

	// 读取数据的任务
	go func() {
		for d:= range data{
			time.Sleep(2*time.Second)
			log.Printf("[接收到了数据,开始处理]：%v", d)
		}

		log.Printf("data chan关闭了，但是我还有清理工作，等我5秒钟")
		time.Sleep(5*time.Second)

		quit <- true
	}()

	// 写入数据
	data <- 1
	time.Sleep(1*time.Second)
	data <-2
	time.Sleep(1*time.Second)
	data <- 3
	time.Sleep(2 * time.Second)
	data <- 4
	log.Printf("发送4")
	data <- 5
	log.Printf("发送5")
	data <- 6
	log.Printf("发送6")
	data <- 7
	log.Printf("发送7")
	data <- 8
	log.Printf("发送8")
	data <- 9
	log.Printf("发送9")

	close(data)
	<- quit
	log.Printf("真正退出了")


}

/*
2021/08/08 09:22:00 [接收到了数据,开始处理]：1
2021/08/08 09:22:03 [接收到了数据,开始处理]：2
2021/08/08 09:22:03 发送4
2021/08/08 09:22:03 发送5
2021/08/08 09:22:03 发送6
2021/08/08 09:22:05 [接收到了数据,开始处理]：3
2021/08/08 09:22:05 发送7
2021/08/08 09:22:07 [接收到了数据,开始处理]：4
2021/08/08 09:22:07 发送8
2021/08/08 09:22:09 [接收到了数据,开始处理]：5
2021/08/08 09:22:09 发送9
2021/08/08 09:22:11 [接收到了数据,开始处理]：6
2021/08/08 09:22:13 [接收到了数据,开始处理]：7
2021/08/08 09:22:15 [接收到了数据,开始处理]：8
2021/08/08 09:22:17 [接收到了数据,开始处理]：9
2021/08/08 09:22:17 data chan关闭了，但是我还有清理工作，等我5秒钟
2021/08/08 09:22:22 真正退出了

Process finished with the exit code 0

*/