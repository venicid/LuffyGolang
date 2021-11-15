package main

import (
	"fmt"
	"sync"

)

// 协程解决方法3：waitGroup，乱序执行
var wtg sync.WaitGroup

func worker(i int)  {
	defer func() {
		fmt.Println(i)
		wtg.Done()
	}()
}

func main()  {

	for i := 0; i < 1000; i++ {
		wtg.Add(1)
		go worker(i)
	}

	wtg.Wait()   // 防止主线程提前结束
	fmt.Println("此处的代码在协程执行完成后，输出")
}

/*
...
996
997
994
999
998
此处的代码在协程执行完成后，输出
*/