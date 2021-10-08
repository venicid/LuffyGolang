package main

import (
	"fmt"
	"time"
)

func main(){

	start := time.Now()

	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)

		close(c)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 1
	}()


	fmt.Println("Blocking...")

	select {
	case <-c:
		fmt.Println("Unblock", time.Since(start))
	case <- ch1:
		fmt.Println("ch1 case")
	case <- ch2:
		fmt.Println("ch2 case")
	//default:
	//	fmt.Println("default go")

	}

}

/*
多个io操作，同时完成，随机选择1个
Blocking...
ch1 case

Blocking...
Unblock 1.005006119s

*/