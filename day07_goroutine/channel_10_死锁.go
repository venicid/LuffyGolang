package main

func main()  {
	ch := make(chan int)

	ch <- 1  // 由于只向main函数中的channel发送，没有goroutinue去接受，导致死锁

}

/*

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /Users/liangshuo276/Learning/HelloGolang/day07_goroutine/channel_10_死锁.go:6 +0x50

*/