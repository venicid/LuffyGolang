package main

import "fmt"

func main(){

	a1 := []int{1,2,3}
	a2 := make([]int,3)  // 指定长度为3
	a3 := make([]int, 4)
	a4 := make([]int, 2)

	copy(a2, a1)
	copy(a3, a1)
	a1[1] = 10
	fmt.Println(a1, a2)
	a2[2] = 20
	fmt.Println(a1, a2)

	fmt.Println(a1, a3)
	fmt.Println(a1, a4)

	/*
	深copy
	[1 10 3] [1 2 3]
	[1 10 3] [1 2 20]
	[1 10 3] [1 2 3 0]
	[1 10 3] [0 0]
	*/
}
