package main

func main() {


	//panic: close of nil channel
	//var c1 chan int
	//close(c1)

	c4:=make(chan int)
	close(c4)

	// panic: close of closed channel
	//c2:=make(chan int)
	//close(c2)
	//close(c2)

	// panic: send on closed channel
	//c3:= make(chan int)
	//close(c3)
	//c3 <- 1
}