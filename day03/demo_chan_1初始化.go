package main

import (
	"log"
	"time"
)

func main()  {

	// 初始化一个int类型的channel
	data := make(chan int)

	//读取数据的任务
	go func() {
		for {
			r := <- data
			log.Printf("[接收到了数据,开始处理]：%v", r)

		}
	}()


	// 写入数据
	data <- 1
	time.Sleep(2*time.Second)
	data <-2
	time.Sleep(2*time.Second)
	close(data)
	log.Println("关闭channel")

	/*
		2021/08/08 09:07:11 [接收到了数据,开始处理]：1
		2021/08/08 09:07:13 [接收到了数据,开始处理]：2
		2021/08/08 09:07:15 关闭channel
	*/
}
