/*
首先，使用channel能达到我们的目的，而且近乎完美
但是，channel在这里，有些大材小用
因为，它被设计出来不仅仅是用作简单的同步处理

使用channel有2个致命的问题
假设有1w或10w个甚至更多的for循环，那就要申请同样数量大小的通道，对内存是不小的开销
*/
package main

import (
	"fmt"
	"sync"
)

var count int

func GetID(wg *sync.WaitGroup)  {
	count = count + 1
	fmt.Println(count)
	wg.Done()
}

func main()  {

	// 协程解决方法3：waitGroup
	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go GetID(&wg)
	}
}

/*
997
735
737
738
739
*/