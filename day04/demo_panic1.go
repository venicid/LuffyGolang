package main

import "fmt"

func defer_func()  {

	defer func() {
		fmt.Println("1")
	}()

	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
	panic("我是panic")
	defer func() {
		fmt.Println("4")
	}()
}

func main()  {
	defer_func()
}

/*
3
2
1
panic: 我是panic
*/