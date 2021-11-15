package main

import (
	"fmt"
	"time"
)

func worker(i int)  {
	fmt.Println(i)
}

func main()  {

	for i := 0; i < 1000; i++ {
		go worker(i)
	}

	// 协程解决方法1：sleep让主线程等待,但是等待多少秒呢
	time.Sleep(time.Millisecond)

}

/*
981
982
983
984
*/