package main

import (
	"fmt"
)

var count int

func worker(ch chan int)  {
	count = count + 1
	fmt.Println(count)
	ch <- count
}

func main()  {

	// 协程解决方法2：channel，顺序执行的goroutinue
	c := make(chan int, 1)

	for i := 0; i < 1000; i++ {
		go worker(c)
		<- c
	}
}

/*
996
997
998
999
1000
*/