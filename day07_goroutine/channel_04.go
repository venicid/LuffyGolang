package main

import (
	"fmt"
	"time"
)

func printer(c chan int)  {
	for  {
		ret := <- c
		fmt.Println("receive", ret)
		if ret == 0{
			break
		}
	}
}

func main(){

	var c = make(chan int)
	go printer(c)

	for i := 10; i >= 0; i-- {
		c <- i
		fmt.Println("send ", i)
		time.Sleep(time.Second)
	}
}