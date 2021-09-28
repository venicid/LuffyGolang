package main

import "fmt"

func foo( i int, c chan  int){
	c <- i
	fmt.Println("send ", i)
}

func main()  {

	c := make(chan int)
	go foo(3, c)

	r := <- c
	fmt.Println("receive ", r)

}

/*
send  3
receive  3

*/
