

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
		for{
			if r,ok := <- data;ok{
				log.Printf("[接收到了数据,开始处理]：%v", r)
			}else{
				log.Printf("没有数据了，chan 关闭了")
			}
			time.Sleep(1*time.Second)
		}

	}()

	// 写入数据
	data <- 1
	time.Sleep(1*time.Second)
	data <-2
	time.Sleep(1*time.Second)

	close(data)
	// 现象是	chan 关闭了没打印
	// 如果加上	time.Sleep(1 * time.Second) 就会打印
	time.Sleep(2*time.Second)

	/*
	2021/08/08 09:15:08 [接收到了数据,开始处理]：1
	2021/08/08 09:15:09 [接收到了数据,开始处理]：2
	2021/08/08 09:15:10 没有数据了，chan 关闭了
	2021/08/08 09:15:11 没有数据了，chan 关闭了
	2021/08/08 09:15:12 没有数据了，chan 关闭了
	*/
}