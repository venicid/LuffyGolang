package main

func main()  {
	ch := make(chan int)
	ch <- 1  // 由于只有香channel发送，但没有goroutine去接受，导致死锁

}
