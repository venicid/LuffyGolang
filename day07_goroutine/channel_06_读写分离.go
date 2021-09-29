package main

import (
	"fmt"
	"time"
)

func send(c chan<- int) {
	fmt.Printf("send: %T\n", c)
	c <- 1
}

func recv(c <-chan int) {
	fmt.Printf("recv: %T\n", c)
	fmt.Println(<-c)
}

func main() {
	c := make(chan int)
	fmt.Printf("%T\n", c)
	go send(c)
	go recv(c)
	time.Sleep(1 * time.Second)
}

/*
chan int
send: chan<- int
recv: <-chan int
1
*/