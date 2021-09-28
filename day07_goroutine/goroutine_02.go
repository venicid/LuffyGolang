package main

import (
	"fmt"
	"time"
)

/*
将running函数并发执行，每隔一秒，打印一次计数器
main的goroutine则等待用户输入
两个行为可以同时进行
*/

func running()  {
	var times int
	for  {
		times += 1
		fmt.Println("count:", times)
		time.Sleep(time.Second * 1)
	}
}

func main1()  {
	// 方法1：启动一个函数
	go running()

	var input string
	fmt.Scanln(&input)

}

func main()  {

	// 方法2：匿名函数
	go func(){
		var times int
		for  {
			times += 1
			fmt.Println("count:", times)
			time.Sleep(time.Second * 1)
		}
	}()

	var input string
	fmt.Scanln(&input)

}