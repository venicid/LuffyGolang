package main

import "time"

func main()  {

	c := make(chan struct{},1)

	// 当有1个缓冲可以用的时候，无需阻塞，可以立即执行
	c <- struct{}{}
	go func() { // 协程1
		//c <- struct{}{}
		time.Sleep(5*time.Second)
		<- c
	}()

	c <- struct{}{}
	go func() {  // 协程1
		//c <- struct{}{}
		time.Sleep(5*time.Second)
		<- c
	}()
}
// channel满了，它会阻塞写；当channel空了，它会阻塞读
// runtime.Goscheud()

/*
无缓冲
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /Users/liangshuo276/Learning/HelloGolang/day07_goroutine/channel_09_阻塞案例.go:10 +0x57
*/