

package main

import (
"log"
"time"
)

func main() {

	// 初始化一个类型为int的chan
	data := make(chan int)

	// 读取数据的任务
	go func() {
		for r:= range data{
			log.Printf("[接收到了数据,开始处理]：%v", r)
		}
		log.Printf("chan 关闭了")
	}()

	// 写入数据
	data <- 1
	time.Sleep(1*time.Second)
	data <-2
	time.Sleep(1*time.Second)

	close(data)
	// 现象是	chan 关闭了没打印
	// 如果加上	time.Sleep(1 * time.Second) 就会打印
	time.Sleep(1*time.Second)
}