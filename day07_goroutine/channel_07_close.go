package main

import "fmt"

func closeDemo(c chan string){

	c <- "a"
	c <- "b"
	close(c)
}

func main()  {

	c := make(chan string)

	go closeDemo(c)

	for i:= range c{
		fmt.Println(i)
	}

}

/*
a
b
*/