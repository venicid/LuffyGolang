package main

import "fmt"

func main()  {

	ch := make(chan int)

	go func() {
		fmt.Println("start child goroutine")
		// 通过通道通知main的goroutine
		ch <- 0
		fmt.Println("end child goroutine")
	}()

	fmt.Println("start main goroutine")
	// 等待匿名goroutine // channel的数据值，如果没有被接受的情况下，是阻塞的
	<- ch
	fmt.Println("all done")

}